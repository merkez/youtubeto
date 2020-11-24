# youtubeto

youtubeto is automated way of saving Youtube Playlists to releases. In this simple hobby project
there are two main files which are  `playlist-list.csv` and `old-playlist-list.csv`. 

`playlist-list.csv` contains, name of the release (- which you would like to have -) and link to Youtube playlist to save as Github release.

## How it works ? 

**Example** `playlist-list.csv`

```csv
cloud_fundamentals_ibm,https://www.youtube.com/playlist?list=PLOspHqNVtKAC-_ZAGresP-i0okHe5FjcJ
cloud_now_ibm,https://www.youtube.com/playlist?list=PLOspHqNVtKAB_5iaF0msu5KQQrOZAdB7k
red_hat_openshit_ibm,https://www.youtube.com/playlist?list=PLOspHqNVtKACZI_rUormOav4THszGESRu
```

Once a commit tagged with semantic versioning, the workflow defined under `.github/workflows/releaseplaylists.yml` will start to run and release defined files in the file.

When workflow successfully finished, all content of `playlist-list.csv` will be appended automatically to `old-playlist-list.csv` file and `playlist-list.csv`will be wiped. 

## How it can be useful ? 

Well, in some cases, Youtube algorithm is showing excellent contents and playlists which you do not would like to lose. 
In those cases, people generally take note the playlists however either they forget in time or playlist became unavailable. 
It creates annoying situations, hence, I have created this mini project, it can be extended more upon requests however currently it is minimal and 
only duty is releasing list of Youtube playlists as Github release assets.  

## Limitations 

There are some limitations regarding to file sizes in releases, according to Github Statement here: 
https://docs.github.com/en/free-pro-team@latest/github/managing-large-files/distributing-large-binaries

Basically, a file size which will be uploaded to releases **should NOT exceeds 2 GB**. However, keep in mind that 
it is per file, **there is NO limitation for overall size of release** :). 
It means that repository will be updated to split files into chunks if size of the file exceeds 2 GB. 
So, in case of 15 GB of playlist, it will be uploaded in 2GB chunks to releases. 

There are some more limitations: 

**Job execution time** - Each job in a workflow can run for up to 6 hours of execution time. If a job reaches this limit, the job is terminated and fails to complete.

**Workflow run time** - Each workflow run is limited to 72 hours. If a workflow run reaches this limit, the workflow run is cancelled.

More details about limmitations on Github Actions: https://docs.github.com/en/free-pro-team@latest/actions/reference/usage-limits-billing-and-administration

It is good to keep in mind the given limitations above. 

**Job execution time**  and **Workflow run time** can be easily fixed if you have your own server. 

If you would like to run Github Actions in your server, there is no limitation regarding to **Job execution time**  and **Workflow run time**.

Check out how to setup Github Actions for your server from here:
 
https://docs.github.com/en/free-pro-team@latest/actions/hosting-your-own-runners/about-self-hosted-runners
