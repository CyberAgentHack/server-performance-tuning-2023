import http from 'k6/http';
import {check, fail} from 'k6';
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';

export const options = {
  thresholds: {
      http_req_duration: ['p(95)<1000'], // 95% のリクエストは 1000ms (1s) 以内に収める
  },
  scenarios: {
      contacts: {
          executor: 'ramping-arrival-rate', // https://k6.io/docs/using-k6/scenarios/executors/ramping-arrival-rate
          exec: 'load_test',

          gracefulStop: '10s',

          preAllocatedVUs: 20,
          stages: [
              { target: 1, duration: '1s' },
              { target: 1, duration: '1s' },
          ],
      },
  },
};

const userIDs = [
  "fc678fcd-c7ba-4735-b0b5-2987fb137e83",
	"ce9ed9b5-3642-4d05-8be4-412881360d4d",
	"cc1bd040-e62f-41c1-abc3-e269f240aefb",
	"bafc4c6f-8c17-440a-b449-071074c48793",
	"3dfd1343-dd72-4a23-859f-af70b1077d92",
	"84019655-24fa-4fc2-8475-d95aa088afb9",
	"f406b15c-5e0c-473c-aabd-5c3e488642fe",
	"2cf08a35-b040-4388-a0be-f01adf8e1202",
	"ee4144c5-d004-45cf-b41f-badd920ccd96",
	"8fff97db-5cc3-497a-b044-bdff4560b0b9",
]

export function fetchEpisodes (offset, limit) {
  const seriesURL =  new URL(`${__ENV.API_BASE_URL}/series`);
  seriesURL.searchParams.append(`limit`, limit);
  seriesURL.searchParams.append(`offset`, `${offset}`);
  
  let res = http.get(seriesURL.toString());
  
  let body = res.json();
  
  const seasonRequests = Array();
  body.series.forEach(series => {
    const url = new URL(`${__ENV.API_BASE_URL}/seasons`);
    url.searchParams.append(`limit`, limit);
    url.searchParams.append(`offset`, `0`);
    url.searchParams.append(`seriesId`, `${series.id}`);
    seasonRequests.push(['GET', url.toString()]);
  });
  
  let seasonRes = http.batch(seasonRequests);
  
  const episodeRequests = Array();
  seasonRes.forEach(
    (res) => {
      const body = res.json();
      if (body === null || body.seasons === null) {
        return;
      }
      body.seasons.forEach((season) => {
        const url = new URL(`${__ENV.API_BASE_URL}/episodes`);
        url.searchParams.append(`limit`, limit);
        url.searchParams.append(`offset`, `0`);
        url.searchParams.append(`seasonId`, `${season.id}`);
        url.searchParams.append(`seriesId`, `${season.seriesId}`);
        episodeRequests.push(['GET', url.toString()]);
      })
    }
  )
  
  const episodes = [];
  const episodesRes = http.batch(episodeRequests);
  episodesRes.forEach((episode) => {
    const body = episode.json();
    episodes.push(...body.episodes);
  })
  
  
  return episodes;
}

export function setup() {
  // const offsets = [0, 50, 100, 150, 200];
  const offsets = [0];
  let episodes = [];

  offsets.forEach((offset) => {
    episodes = episodes.concat(fetchEpisodes(offset, 50));
  })
  console.log(episodes.length);
  return { episodes: episodes };
}

export function load_test(data) {
  const episodes = data.episodes;
  
  console.log(episodes[0]);
  userIDs.forEach((userID) => {
    const viewingHistoryRequests = Array();
    const episodeIDs = Array();

    episodes.forEach((episode) => {
      episodeIDs.push(episode.id);
    });

    const url = new URL(`${__ENV.API_BASE_URL}/viewingHistories`); 
    const params = {
      headers: {
        "userId": userID,
      },
    }
    url.searchParams.append(`episodeIds`, episodeIDs.join(","));
    viewingHistoryRequests.push(['GET', url.toString()]);
    const res = http.get(url.toString(), params);
    check_status_ok(res);
  })

}

function check_status_ok(res) {
  const result = check(res, {
    'is status OK': (r) => r.status === 200,
  }); 
  if (!result) {
    fail(`status is not 200. response: ${JSON.stringify(res)}`);
  }
}

