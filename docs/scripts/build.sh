#!/usr/bin/env bash

rm -rf site/
docker run --rm -it -v ${PWD}:/docs squidfunk/mkdocs-material build
