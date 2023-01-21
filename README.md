# KSTORAGE Processing

That is the App that allows to sync two volumes

## Run app

go build -o kstorage-processing
Run with two params `fromVolume` as source volume and `toVolume` as destination volume

./kstorage-processing --fromVolume=/Users/dmitriykrylov/Documents/work/projects/go/kstorage-processing/dir1 --toVolume=/Users/dmitriykrylov/Documents/work/projects/go/kstorage-processing/dir2

If `disableDeletion` flag is on, than the App will not process deletion of the old(removed from the source storage) files

### Details

How does it work in details: hash, .volumes, etc.