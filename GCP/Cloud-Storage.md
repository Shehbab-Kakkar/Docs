what is the maximum size of an object in cloud storage in gcp

The maximum size of an individual object that can be stored in Google Cloud Storage (GCP Cloud Storage) is 5 tebibytes (5 TiB)[1][2][3][4][5]. Objects larger than 5 TiB cannot be uploaded or stored as a single file in GCP Cloud Storage[1]. 

### Object Size Limit Details

- A single object (file) can be up to **5 TiB** in size[1][2][3][4][5].
- If a multipart upload is used, the assembled object's total size must also not exceed 5 TiB[3].
- Buckets themselves do not have a total maximum storage limit; only the per-object (file) limit applies[2].

This 5 TiB object size limitation is consistent across most official documentation sources for GCP Cloud Storage as of September 2025[1][4][5].

Citations:
[1] Known Limitations | Cloud Storage Transfer Service ... https://cloud.google.com/storage-transfer/docs/known-limitations-transfer
[2] What is the maximum size of a bucket on Google Cloud ... https://stackoverflow.com/questions/26908695/what-is-the-maximum-size-of-a-bucket-on-google-cloud-storage-when-should-i-use
[3] GCP Cloud storage - DevOps Blogs https://dhirajubhalkar.hashnode.dev/gcp-cloud-storage
[4] Cloud Storage - Google Cloud Platform Console Help https://support.google.com/cloud/answer/6250993?hl=en
[5] Google Cloud Storage And Database Services Rundown https://k21academy.com/google-cloud/google-cloud-storage-and-database/
[6] Quotas & limits | Cloud Storage https://cloud.google.com/storage/quotas
[7] Maximum File size limit https://groups.google.com/g/gs-discussion/c/DkrQ_N09jl8
[8] Storage limits in Google Cloud https://docs.netapp.com/us-en/cloud-volumes-ontap-96-relnotes/reference-limits-gcp.html
[9] Storage limits for Cloud Volumes ONTAP in Google Cloud https://docs.netapp.com/us-en/cloud-volumes-ontap-relnotes/reference-limits-gcp.html
