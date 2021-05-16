const { URL } = require('url')
const fs = require('fs')

const rawUrls = fs.readFileSync('addUrl.txt', 'utf8')
const uBlacklist = fs.readFileSync('uBlacklist.txt', 'utf8')

const string2array = string => string.trim().split('\n')

const handleHostname = (list = []) =>
  list.reduce((result, url) => {
    try {
      const x = new URL(url).hostname.replace('www.', '')
      return [...result, `*://*.${x}/*`]
    } catch (e) {
      return result
    }
  }, [])

const result = new Set([
  ...string2array(uBlacklist),
  ...handleHostname(string2array(rawUrls))
])

fs.unlinkSync('uBlacklist.txt')
const fileStream = fs.createWriteStream('uBlacklist.txt', { flags: 'a' })

result.forEach(url => {
  fileStream.write(url + '\n')
})
