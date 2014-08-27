# Shacl

Life is too short to keep seeing duplicate postings on Craigslist.

_Systemizing Apartments and Housing for Craigslist_ not _SHAing Craiglist_, or is it the other way around?

This is a general purpose RSS reader for Craigslist.

It takes the feed, SHA1s the title and body. You mark the item as read or not.

Once an item is read anytime another item with the matching body or title is filtered out.


### TODO

* x Download the RSS feed
* x Parse the feed
* x Store a hash of the title and body
* x UI for read/unread
* x store read items
* x filter out read items (via hash)
* x Do not redisplay articles which match the title or body hash.
* x Web UI
* x Dameon to auto run
* Pretty UI
* Gracefully refresh new listings via EventSource
* Email/Text new listings
* Configure the RSS url
* Add accounts
* Profit!
