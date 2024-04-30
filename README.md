# HALIE SYNC TOOL #


An intermediary tool used by various other tools, that syncs data and assets between local, staging and live.

Install the tool using Homebrew.
To do so, you'll need to gain access to our Github account. Get the Personal Token from the Github entry notes in Bitwarden and use it to run the following command:

    export HOMEBREW_GITHUB_API_TOKEN=<--TOKEN HERE-->


Now you can 'tap' our Homebrew tools by running:

    brew tap ivystreetweb/homebrew-tools


Allowing you to install the LSET tool:

    brew install hsync

You will also need to retrieve the halie_sync SSH key, which can be found in Bitwarden.

This will install our tool, along with the Go binaries.

In case you need to use this tool by itself, you can run it manually on the command line using flags.

Go to the root folder where your assets are stored or where you'd like them to be, then run the command with flags, using this structure:

    hsync sync --flags


Look up the man file for available flags. The flags will allow you to define whether you are pushing/pulling, where, and what type of content.


For example, this command will pull down the data JSON files and masterplan files from your specified live server, to the folder you're running the command from:

    hsync sync --pull --live --data --masterplan

Another example, this will push media files and image/PDF assets to the staging environment:

    hsync sync --push --staging --assets --media

You can also use the --all flag to sync everything.