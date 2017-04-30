# Low Frequency Trader #

## Usage ##
There are three primary modes this trader can be run in: historical mode; live-with-fake-money; and live-with-real-money. The first mode, historical, is intended to be your starting point for developing a theory. Once you've tweaked your algorithm enough, move on to live-with-fake-money for a few weeks. This will feed live data from a live brokerage to your algorithm. You can see if you would have made money over the investigation period. Finally, move on to running with live data and real money when you're ready for the big leagues.

### Historical Mode ###
`trader historical 1000 coindesk coindesk_history.csv --graphs`

`1000` - the principal investment, in USD.
`coindesk` - the name of the brokerage and gives information about how to parse the provided csv. It also contains information about applicable fees that'll get ya.
`coindesk_history.csv` - The history as a CSV, formatting dependent upon the brokerage.
`--graphs` - [Optional] produce pretty pictures about when buys or sells were triggered.

### Live without money ###
`trader nomoney-live 1000 coindesk --graphs`

### Live with real money ###
`trader live coindesk --graphs`
