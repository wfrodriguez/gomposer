local fn = require('funcs')

require('config.go')
require('config.web')

keyOpts = {noremap = true, silent = true}
fn.nmap('<F5>', ':GoFmt<CR>', keyOpts)
