# quick-restock

This is a tiny web server which I wrote while exploring using Go for personal projects. I use Go at Google for my day job but I've tended to use TypeScript for personal projects for years now.

All this application does is redirect you to a configured URL based on the path you request. This is used as part of an automation I have setup on my phone using Apple Shortcuts. Triggering the shortcut immediately gives me a camera view to scan barcodes on items, then by communicating with this server I am deep-linked into the corresponding app for repurchasing the item. It's just a quick way to restock items I've run out of.
