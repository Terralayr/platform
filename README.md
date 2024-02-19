# platform

This repo is designed to help run the terralayr platform locally.

## Setup

This will attempt to use local versions of services, and assumes they are stored in directories 
one level up from this repo. e.g. the relative path to the `virtual-assets` repo should be 
`../virtual-assets`.

## Running


```bash
make up

```

Should bring up the relevant services, apply migrations, and fill out some test data.

