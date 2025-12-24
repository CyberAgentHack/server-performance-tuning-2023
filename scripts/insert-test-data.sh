#!/bin/bash

set -e

echo "Inserting test data into MySQL..."

# genres
docker exec -i wsperf-mysql mysql -u root wsperf << 'EOF'
INSERT INTO genres (genreID, displayName) VALUES
('976-755', 'アニメ'),
('878-285', 'ドラマ'),
('51-872', 'バラエティ'),
('816-680', '映画'),
('719-306', '格闘技');
EOF

echo "✓ Inserted 5 genres"

# series
docker exec -i wsperf-mysql mysql -u root wsperf << 'EOF'
INSERT INTO series (seriesID, displayName, description, imageURL, genreID) VALUES
('374-745', 'ONE PIECE', 'これは作品 ONE PIECE の説明文です。様々なジャンルの面白いコンテンツをお楽しみください。', 'https://image.p-c2-x.abema-tv.com/image/series/374-745', '976-755'),
('77-257', '名探偵コナン', 'これは作品 名探偵コナン の説明文です。様々なジャンルの面白いコンテンツをお楽しみください。', 'https://image.p-c2-x.abema-tv.com/image/series/77-257', '878-285'),
('937-848', 'ドラゴンボール', 'これは作品 ドラゴンボール の説明文です。様々なジャンルの面白いコンテンツをお楽しみください。', 'https://image.p-c2-x.abema-tv.com/image/series/937-848', '51-872'),
('703-400', 'NARUTO', 'これは作品 NARUTO の説明文です。様々なジャンルの面白いコンテンツをお楽しみください。', 'https://image.p-c2-x.abema-tv.com/image/series/703-400', '816-680'),
('565-598', '進撃の巨人', 'これは作品 進撃の巨人 の説明文です。様々なジャンルの面白いコンテンツをお楽しみください。', 'https://image.p-c2-x.abema-tv.com/image/series/565-598', '719-306');
EOF

echo "✓ Inserted 5 series"

# seasons
docker exec -i wsperf-mysql mysql -u root wsperf << 'EOF'
INSERT INTO seasons (seasonID, seriesID, displayName, imageURL, displayOrder) VALUES
('374-745_s1', '374-745', 'シーズン1', 'https://image.p-c2-x.abema-tv.com/image/series/374-745/season1', 1),
('374-745_s2', '374-745', 'シーズン2', 'https://image.p-c2-x.abema-tv.com/image/series/374-745/season2', 2),
('77-257_s1', '77-257', 'シーズン1', 'https://image.p-c2-x.abema-tv.com/image/series/77-257/season1', 1);
EOF

echo "✓ Inserted 3 seasons"

# episodes
docker exec -i wsperf-mysql mysql -u root wsperf << 'EOF'
INSERT INTO episodes (episodeID, seasonID, seriesID, displayName, description, imageURL, displayOrder) VALUES
('374-745_s1_e1', '374-745_s1', '374-745', '第1話', 'ONE PIECE 第1話の説明', 'https://image.p-c2-x.abema-tv.com/image/episode/374-745_s1_e1', 1),
('374-745_s1_e2', '374-745_s1', '374-745', '第2話', 'ONE PIECE 第2話の説明', 'https://image.p-c2-x.abema-tv.com/image/episode/374-745_s1_e2', 2),
('77-257_s1_e1', '77-257_s1', '77-257', '第1話', '名探偵コナン 第1話の説明', 'https://image.p-c2-x.abema-tv.com/image/episode/77-257_s1_e1', 1);
EOF

echo "✓ Inserted 3 episodes"

echo ""
echo "Test data inserted successfully!"
echo "Total: 5 genres, 5 series, 3 seasons, 3 episodes"
