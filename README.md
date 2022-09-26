# charter-challenge-for-fair-voting-bot

A bot that records the amount donated so far for the Charter Challenge for Fair Voting (https://www.charterchallenge.ca/) into a public BigQuery dataset each day.

## Querying data

The data is available in a public dataset that is updated hourly. Example query:

```sql
SELECT
  *
FROM
  `public-datasets-363301.charter_challenge_bot.donation_statuses`
ORDER BY
  date DESC
LIMIT
  1000
```

Any GCP principal authenticated can query this public dataset.

![image](https://user-images.githubusercontent.com/7719209/191647190-3bcb7f7c-ed19-49b9-881a-7e4ebe685864.png)

![image](https://user-images.githubusercontent.com/7719209/191645928-6068d947-c777-41ab-ae94-0d38d3c56110.png)

The data is stored per hour with the string from the website stored as is, but you can use BigQuery SQL to further process the data. For example:

```sql
SELECT
  # Truncate timestamps to just their day
  DATE(TIMESTAMP_TRUNC(date, DAY)) AS date,
  # Convert string from web page to NUMERIC and get a summary
  # of how much had been donated by the end of that day. Format
  # as currency at the end to display nicely.
  CONCAT('$',FORMAT("%'.2f", MAX(CAST(REGEXP_REPLACE(so_far, '[^.0-9 ]', '') AS NUMERIC)))) AS donated_so_far
FROM
  `public-datasets-363301.charter_challenge_bot.donation_statuses`
GROUP BY
  date
ORDER BY
  date DESC
```

![image](https://user-images.githubusercontent.com/7719209/192045133-2591727c-bfc0-4998-a291-ce67ab2699db.png)

You can use BigQuery [window function calls](https://cloud.google.com/bigquery/docs/reference/standard-sql/window-function-calls) to look for significant parts of the data, like during which hours the donations were made, instead of just seeing a running total by querying the raw data:

```sql
WITH
  t1 AS (
  SELECT
    TIMESTAMP_TRUNC(date, HOUR) AS hour,
    CAST(REGEXP_REPLACE(so_far, '[^.0-9 ]', '') AS NUMERIC) AS donated_so_far
  FROM
    `public-datasets-363301.charter_challenge_bot.donation_statuses`
  ORDER BY
    hour DESC ),
  t2 AS (
  SELECT
    hour,
    donated_so_far - LAG(donated_so_far, 1) OVER (ORDER BY hour) AS donated_during_hour
  FROM
    t1)
SELECT
  hour,
  CONCAT('$',FORMAT("%'.2f", donated_during_hour)) AS donated_during_hour
FROM
  t2
WHERE
  donated_during_hour > 0
ORDER BY
  hour DESC
```

![image](https://user-images.githubusercontent.com/7719209/192328899-573b1992-6790-42b7-b176-061019bfc0e6.png)

## Credits

Example used for guidance: https://github.com/apache/openwhisk-runtime-go/tree/master/examples/module-main
