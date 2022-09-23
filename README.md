# charter-challenge-for-fair-voting-bot

A bot that records the amount donated so far for the Charter Challenge for Fair Voting (https://www.charterchallenge.ca/) into a public BigQuery dataset each day.

## Reading data

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
  # of how much had been donated by the end of that day
  MAX(CAST(REGEXP_REPLACE(so_far, '[^.0-9 ]', '') AS NUMERIC)) AS so_far
FROM
  `public-datasets-363301.charter_challenge_bot.donation_statuses`
GROUP BY
  date
ORDER BY
  date DESC
```

![image](https://user-images.githubusercontent.com/7719209/192031064-dfaee263-4286-4433-a90a-3fb069b59def.png)

## Credits

Example used for guidance: https://github.com/apache/openwhisk-runtime-go/tree/master/examples/module-main
