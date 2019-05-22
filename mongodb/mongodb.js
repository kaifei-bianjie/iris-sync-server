// create collections
db.createCollection("account");
db.createCollection("block");
db.createCollection("stake_role_candidate");
db.createCollection("stake_role_delegator");
db.createCollection("sync_task");
db.createCollection("tx_common");
db.createCollection("validator_up_time");
db.createCollection("tx_gas");

// create index
db.account.createIndex({"address": 1}, {"unique": true});
db.block.createIndex({"height": -1}, {"unique": true});

db.stake_role_candidate.createIndex({"address": 1}, {"unique": true});
db.stake_role_candidate.createIndex({"pub_key": 1});

db.stake_role_delegator.createIndex({"validator_addr": 1});
db.stake_role_delegator.createIndex({"address": 1});
db.stake_role_delegator.createIndex({"address": 1, "validator_addr": 1}, {"unique": true});

db.sync_task.createIndex({"chain_id": 1}, {"unique": true});

db.tx_common.createIndex({"height": -1});
db.tx_common.createIndex({"time": -1});
db.tx_common.createIndex({"tx_hash": 1}, {"unique": true});
db.tx_common.createIndex({"from": 1});
db.tx_common.createIndex({"to": 1});
db.tx_common.createIndex({"type": 1});
db.tx_common.createIndex({"status": 1});

db.validator_up_time.createIndex({"val_address": 1}, {"unique": true});

db.tx_gas.createIndex({"tx_type": 1}, {"unique": true});

// drop collection
// db.account.drop();
// db.block.drop();
// db.stake_role_candidate.drop();
// db.stake_role_delegator.drop();
// db.sync_task.drop();
// db.tx_common.drop();
// db.validator_up_time.drop();
// db.tx_gas.drop();
//
// remove collection data
// db.account.remove({});
// db.block.remove({});
// db.stake_role_candidate.remove({});
// db.stake_role_delegator.remove({});
// db.sync_task.remove({});
// db.tx_common.remove({});
// db.validator_up_time.remove({});
// db.tx_gas.remove({});










