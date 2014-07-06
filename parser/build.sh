#!/usr/bin/env bash

go tool yacc -o parser.go -p Ruby parser.y
