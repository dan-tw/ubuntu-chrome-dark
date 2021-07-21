# ubuntu-chrome-dark
Change the internal pages on chrome to use dark mode

## How to use

For this to work your chrome theme must be set to `classic` and `dark-mode` must be enabled in Ubuntu.

Clone the repo:

```
git clone https://github.com/dan-tw/ubuntu-chrome-dark.git
```

CD into the repo directory and run the binary (must use `sudo`). This will update the necessary file to add `dark-mode` to your chrome browser

```
cd ubuntu-chrome-dark
sudo ./ubuntu-chrome-dark
```

### Notes

Running the binary multiple times will not matter so if you wanted to set it up to run on login you could (this will catch when chrome updates and `dark-mode` may disappear!)
