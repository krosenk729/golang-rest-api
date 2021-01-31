#!/usr/bin/env sh
set -e

initiateReplicasetAndCompatibilityVersion() {
  sleep 10 && mongo admin --eval 'rs.initiate({ "_id": "rs1", "members" : [ {"_id": 1, "host": "localhost:27017"} ]}); db.adminCommand( { setFeatureCompatibilityVersion: "4.4" } );'
}

initiateReplicasetAndCompatibilityVersion &

mongod --replSet rs1 --bind_ip 0.0.0.0
