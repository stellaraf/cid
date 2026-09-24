# CI runner routing and retention

Linux CI runs on disposable isolated public AWS CodeBuild workers. Native macOS and Windows jobs retain their platforms. CI build artifacts use the private public-repository S3 bucket, scoped by repository, run, attempt, and commit. Build artifacts expire after 30 days; dependency caches expire after 14 days. Published GitHub Release assets keep their release lifecycle.
