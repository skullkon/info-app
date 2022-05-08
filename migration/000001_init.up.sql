CREATE TABLE IF NOT EXISTS info
(
    id             UUID,
    ip             String,
    type           String,
    os             String,
    osVersion      String,
    browser        String,
    browserVersion String,
    brand          String,
    model          String,
    resolution     String,
    time           DateTime
) ENGINE = MergeTree()
      PARTITION BY toYYYYMM(time)
      ORDER BY (id, browser, brand, os);

