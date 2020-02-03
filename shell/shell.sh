#!/bin/bash
redis-cli keys W*| xargs -r -t -n1 redis-cli  del