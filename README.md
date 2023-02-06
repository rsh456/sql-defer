## Defer usage example in GO
- Script for create activities sqlite db
- Connect to DB usign sqlite3 driver
- Use defer for closing connections, first in -last out ordering 
(The last deferred function is the first one to be invoked)
