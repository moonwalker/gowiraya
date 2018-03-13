# GoWiraya

Go library for sending SMS with Wiraya.

Note that the ip needs to be whitelisted at Wiraya.

They also have some bugs
1. They do not support HTTP/2 over SSL
2. They do see headers as case sensitive (RESOLVED but not in production yet)