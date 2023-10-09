#!/bin/sh
# Retrieve the .onion domain and update the environment variable
onion_domain=$(cat /var/lib/tor/hidden_service/hostname)
export ONION_DOMAIN="$onion_domain"