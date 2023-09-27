function abbr(a, w)
	vim.cmd("iabbrev " .. a .. " " .. w)
end

local fn = require("functions")

vim.opt.expandtab = false
vim.opt.shiftwidth = 4
vim.opt.tabstop = 4
abbr("ee", ":=")
fn.nmap("<F5>", ":VimuxRunLastCommand<CR>")
