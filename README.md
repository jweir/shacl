# Shacl

Life is too short to keep seeing duplicate postings on Craigslist.

_Systemizing Apartments and Housing for Craigslist_ not _SHAing Craiglist_, or is it the other way around?

This is a general purpose RSS reader for Craigslist.

It takes the feed, SHA1s the title and body. You mark the item as read or not.

Once an item is read anytime another item with the matching body or title is filtered out.


### TODO

* Download the RSS feed
* Parse the feed
* Store a hash of the title and body
* UI for read/unread
* store read items
* filter out read items (via hash)
* Do not redisplay articles which match the title or body hash.
* Web UI
* Dameon to auto run

Store the feed (need to accumlate it, not just store in memory)
Have a log file of hashes, no need for any sort of db, only hashes which have been marked read will be written.

Disk looks like
  exec
  .feed
  .read

Prune the feed occasionaly


