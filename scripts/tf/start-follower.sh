#!/bin/bash
nohup ./follower -coordinator=$(cat coordinator.address) -label=$(cat follower.label) -time=3h &