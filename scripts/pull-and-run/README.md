# pull-and-run

![Bash](https://img.shields.io/badge/Bash-4EAA25?style=flat&logo=gnubash&logoColor=white)

Automated script that pulls the latest Docker images and redeploys the services by running the `Makefile` defined in `/opt/hauslab/services`.

## Cron job

It runs as cron job every 10 minutes through the following entry in `crontab`:

```bash
*/10 * * * * /opt/hauslab/scripts/pull-and-run/pull-and-run.sh >> /opt/hauslab/scripts/pull-and-run/pull-and-run.log 2>&1
```

To modify the entry, run:

```bash
crontab -e
```

The output of the script is stored in `/opt/hauslab/scripts/pull-and-run/pull-and-run.log`, which is weekly deleted by another cron job defined in `/etc/cron.weekly`.
