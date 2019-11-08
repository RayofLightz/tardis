# Tardis
Tardis is an DVR/IP surveillance camera honeypot.

## Why tardis
Tardis was built with the intention to harvest credentials being used from automated IP camera scanners. 

Tardis was given its name because I came up with the idea for this project while watching Dr. Who 
and because the tardis looks deceptive on the outside like most honeypots.

## Install

`go get github.com/RayofLightz/tardis`

## Configuration

Configuration values can be found in `config/config.json` file.

## Logs
Tardis creates logs when fake sucess pages are generated and when authentication attempts are made against the server.
