# buchladen

---

- [DESIGN OVERVIEW](#design-overview)

## DESIGN OVERVIEW

This is an experimental design of a microservice-oriented software, the
main goal of which is to interact with users and keep track of items
they sell.

We need to have an API to store various users of our service as well as
manipulate such data:

```
                    _____
+---------+        /     \
|  Users  | <---> ( MySQL )
|   API   |       (\_____/)
+---------+       (_______)
```

[OAuth](https://www.rfc-editor.org/rfc/rfc6749) authorisation framework
shall be used for user authentication. Since we’ve taken the
microservice-oriented approach, we shall create a new NoSQL
database—[Apache Cassandra](https://cassandra.apache.org) in this
case—to be used in conjunction with our OAuth API to store access tokens
and whatnot:

```
  _______
 /       \
( Cassan- )
(   dra   )
(\_______/)
(_________)
(_________)

     ^
     |
     |
     v
+---------+
|  OAuth  |
|   API   |
+---------+
```

We also need to store the list of items users are selling at the moment
or have already sold:

```
+---------+
|  Items  |
|   API   |
+---------+
     ^
     |
     |
     v
   _____
  /     \
 ( MySQL )
 (\_____/)
 (_______)
```

It may be a good idea for MySQL to synchronise against the
[Elasticsearch](https://www.elastic.co/elasticsearch) RESTful search
engine, with our items API taking input from both MySQL and ELS:

```
         +---------+
     +-> |  Items  |
     |   |   API   | <-+
     |   +---------+   |
     |                 |
     v                 |
   _____             _____
  /     \           /     \
 ( MySQL ) ------> (  ELS  )
 (\_____/)         (\_____/)
 (_______)         (_______)
```

Items API cannot be of much help to us if it doesn’t know what token is
given to which user, so taking input from the OAuth API is paramount:

```
+---------+        +---------+
|  Items  | <----- |  OAuth  |
|   API   |        |   API   |
+---------+        +---------+
```

This is pretty much it. Here is the final overview of the design with
all the elements being in one place:

```
                       _______
                      /       \
                     ( Cassan- )
                     (   dra   )
                     (\_______/)
                     (_________)
                          ^
                          |
                     +---------+
                     |  Oauth  |
              +----- |   API   | <---+
              |      +---------+     |
              v                      |
         +---------+                 |              _____
     +-> |  Items  |            +---------+        /     \
     |   |   API   | <-+        |  Users  | <---> ( MySQL )
     |   +---------+   |        |   API   |       (\_____/)
     |                 |        +---------+       (_______)
     v                 |
   _____             _____
  /     \           /     \
 ( MySQL ) ------> (  ELS  )
 (\_____/)         (\_____/)
 (_______)         (_______)
```
