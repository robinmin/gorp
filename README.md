# Fix for Gorp #

This is a fork on [https://github.com/coopernurse/gorp](https://github.com/coopernurse/gorp), for the details useage, please refer to the original place.

## Enhancement ##
  - Add Support on SQL Server. So far, My prefference golang driver for SQL server is [https://github.com/mattn/go-adodb](https://github.com/mattn/go-adodb) -- (2014-04-14).

## Todo ##
  - Support different table creation syntax;
  - Add ping/pong test after the connection established, otherwise, the system cannot tell the connect is valid or not;