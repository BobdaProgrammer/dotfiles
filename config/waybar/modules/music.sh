#!/bin/sh
name=$(playerctl metadata --format='{{ title }}')
class=$(playerctl metadata --player=spotify --format '{{lc(status)}}')
icon=""

if [[ $class == "playing" || $class == "paused" ]]; then
	echo "$name $icon"
else
	echo "$name"
fi


