// Create a client
// Set a key-value pair in Redis
// Get the value of the key from Redis


## Redis data types

Redis stores data as keys and values. The key is usually a string; the value can be a string, hash, list, set, or another supported type.

Think of the key as a variable name and the data type as the kind of value it holds.

These examples run in redis-cli.
1. String — one value
A Redis string can hold text, a number, or serialized data such as JSON.
SET user:101:name "Gaurav"
GET user:101:name
# "Gaurav"

The : is just a naming convention. Redis doesn’t treat it as a folder.

Numbers can be incremented directly:
SET page:views 100
INCR page:views
# 101

INCRBY page:views 10
# 111

You can also give a key an expiration:
SET session:abc "user-101" EX 3600

This key disappears after 3,600 seconds.

Useful for: cached responses, counters, session tokens.

A JSON string is still a string to Redis:
SET user:101 '{"name":"Gaurav","city":"Jersey City"}'

With string commands, you retrieve or replace the whole value. Redis doesn’t interpret its JSON fields.
2. Hash — fields inside one key
A hash holds field-value pairs. Think of it as a small object or dictionary.
HSET user:102 name "Alice" city "New York" age 30

Conceptually:
user:102 → {name: "Alice", city: "New York", age: "30"}

Read one field:
HGET user:102 name
# "Alice"

Read all fields:
HGETALL user:102

Update a field without replacing the entire object:
HSET user:102 city "Boston"
HINCRBY user:102 age 1

Useful for: user profiles, product details, small records.

String vs. hash: use a string when you read/write the whole value; use a hash when you need to access individual fields.

3. List — ordered items, duplicates allowed
A list preserves item order. You add or remove items at either end.
RPUSH tasks "build" "test" "deploy"

The list now contains:
["build", "test", "deploy"]

- LPUSH: add to the left.
- RPUSH: add to the right.
- LPOP: remove from the left.
- RPOP: remove from the right.
LPOP tasks
# "build"

LRANGE tasks 0 -1
# "test", "deploy"

0 -1 means “from the first item through the last.”

Duplicates are allowed:
RPUSH tasks "test"
# ["test", "deploy", "test"]

Useful for: recent activity and simple queues.

##
```
For example, RPUSH to enqueue and LPOP to dequeue gives FIFO behavior. A production job queue also needs a strategy for failed workers and retries.
```

6. Stream — an append-only sequence of events
A stream stores events with IDs. Unlike popping from a list, reading an event doesn’t remove it.
XADD orders:events "*" order_id 501 status "created"

"*" asks Redis to generate an event ID.

Add another event:
XADD orders:events "*" order_id 501 status "paid"

Read the events:
XRANGE orders:events - +

Here, - + means “from the earliest through the latest.”

4. Set — unique items, no guaranteed order
A set answers: “Is this item a member?”
SADD online:users "alice" "bob" "alice"

Only two members exist because duplicates are ignored.
SCARD online:users
# 2

SISMEMBER online:users "alice"
# 1 — yes

SISMEMBER online:users "charlie"
# 0 — no

SMEMBERS online:users
# "alice", "bob" — order can vary

Sets also support mathematical operations:
SADD team:platform "alice" "bob"
SADD team:security "bob" "charlie"

SINTER team:platform team:security
# "bob"

Useful for: membership, tags, deduplication, finding overlap between groups.


5. Sorted set — unique items ordered by score
A sorted set gives each member a numeric score. Redis orders members by that score.
ZADD leaderboard 120 "alice" 90 "bob" 150 "charlie"

Read from lowest score to highest:
ZRANGE leaderboard 0 -1 WITHSCORES
# bob 90, alice 120, charlie 150

Read highest first:
ZRANGE leaderboard 0 -1 REV WITHSCORES
# charlie 150, alice 120, bob 90

Increase someone’s score:
ZINCRBY leaderboard 40 "bob"
# bob now has 130

Members are unique, but different members can have the same score. Adding an existing member updates its score.

Useful for: leaderboards, rankings, scheduling items by timestamp.

=====

Commands that work across types 
TYPE user:102
# hash

EXISTS user:102
# 1

EXPIRE user:102 60
# Expire the key in 60 seconds

TTL user:102
# Remaining lifetime in seconds

DEL user:102
# Delete the key
