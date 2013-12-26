#!/bin/bash

DOCKER='docker'

$DOCKER build -t test/docker-chrome-selenium docker-chrome-selenium
$DOCKER build -t test/selenium-client selenium-client
$DOCKER run --privileged -d -name selenium-chrome test/docker-chrome-selenium
$DOCKER run -t -i -link selenium-chrome:chrome test/selenium-client

