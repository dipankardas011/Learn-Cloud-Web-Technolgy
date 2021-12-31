# Difference b/w the formal flow & Async & Difer

To summarize:

* async and defer both instruct the browser to download the script(s) in a separate thread, while the rest of the page (the DOM, etc.) is downloading, so the page loading is not blocked during the fetch process.
* scripts with an async attribute will execute as soon as the download is complete. This blocks the page and does not guarantee any specific execution order.
* scripts with a defer attribute will load in the order they are in and will only execute once everything has finished loading.
* If your scripts should be run immediately and they don't have any dependencies, then use async.
* If your scripts need to wait for parsing and depend on other scripts and/or the DOM being in place, load them using defer and put their corresponding `<script>` elements in the order you want the browser to execute them.
