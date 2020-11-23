# Mixtape

### About

A music-centered social media app that allows users to create and authenticate their account, create posts, share songs and albums and artists, add friends, and view a timeline of their friend's posts and shared music.

### Tools, Technologies, and Resources
The app currently utilizes the following tools and technologies:
 - the Go programming language
 - a MySQL database for storing users, profiles, and posts
 - Docker for microservice containers
 - a Go wrapper for working with Spotify's API (found here: https://pkg.go.dev/github.com/zmb3/spotify)
 - Amazon Neptune for the friends portion of the application which uses a graph database
 - AWS Educate and EC2 (Elastic Cloud Compute) to deploy and host the application on a Virtual Machine.

### Notes
In the works:
 - Implementing Spotify endpoints
 - Integrating front-end
 - Setting up Amazon services
