# platform

This repo is designed to help run the terralayr platform locally.

## Setup

This will attempt to use local versions of services, and assumes they are stored in directories 
parallel to the location of this repo. e.g. the relative path to the `virtual-assets` repo should
be `../virtual-assets`.

## Running

```bash
make up
```

Should bring up the relevant services, apply migrations, and fill out some test data.

You should then be able to hit the gateway service at `localhost:8000`