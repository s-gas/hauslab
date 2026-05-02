# pull-and-run

![Bash](https://img.shields.io/badge/Bash-4EAA25?style=flat&logo=gnubash&logoColor=white)

Script that pulls the latest commit and redeploys the services by running the `Makefile` defined in `~/hauslab/services`.

## Cron job

It runs as cron job every 10 minutes through the following entry in `crontab`:

```bash
*/10 * * * * /home/s-gas/hauslab/scripts/pull-and-run/pull-and-run.sh >> /home/s-gas/hauslab/scripts/pull-and-run/pull-and-run.log 2>&1
```

To modify the entry, run:

```bash
crontab -e
```
