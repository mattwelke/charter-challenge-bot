Some steps I can run to manually create new builds and deploy them after updates.

1. Build image

   ```
   docker build -t northamerica-northeast2-docker.pkg.dev/public-datasets-363301/charter-challenge-scraper/scraper .
   ```

1. Push image

   ```
   docker push northamerica-northeast2-docker.pkg.dev/public-datasets-363301/charter-challenge-scraper/scraper
   ```

1. Deploy job

   ```
   gcloud config set project public-datasets-363301
   ```

   ```
   gcloud run jobs deploy charter-challenge-scraper \
     --image northamerica-northeast2-docker.pkg.dev/public-datasets-363301/charter-challenge-scraper/scraper \
     --region northamerica-northeast2 \
     --service-account charter-challenge-bot@public-datasets-363301.iam.gserviceaccount.com
   ```

The Cloud Scheduler trigger has already been set up and doesn't change, so every time there is an update to the code, I just need to update the Cloud Run job by completing the steps above.
