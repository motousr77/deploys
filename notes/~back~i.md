* backup infrastructure (terraform/other tools/custom format)
* backup cluster state (etcd/other db)
* backup persistent data (tarball/zipball)
* regular backups / backup scheduler (job/chronejob/custom shceduler)
* backup by event (event handler/job starter)
* backup and package secrets to some of private cloud
* state changes history viewer
* elastic cluster analizer tool
_ 

~~~
known implementations:
reshifter -  >> https://github.com/mhausenblas/reshifter

velero - backaup cluster state >> https://velero.io/docs/v1.2.0/how-velero-works/
https://github.com/vmware-tanzu/velero

runcher - import/export infrastructure >> https://github.com/rancher/rke/wiki/Quick-Start-Guide
https://rancher.com/docs/rancher/v2.x/en/quick-start-guide/deployment/quickstart-manual-setup/
https://rancher.com/docs/rancher/v1.6/en/quick-start-guide/
