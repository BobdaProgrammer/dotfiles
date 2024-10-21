-- Ensure lazy.nvim is installed
local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
  vim.fn.system({
    "git",
    "clone",
    "--filter=blob:none",
    "--branch=stable", -- latest stable release
    "https://github.com/folke/lazy.nvim.git",
    lazypath,
  })
end

vim.opt.rtp:prepend(lazypath)
require("lazy").setup({
  { 
    "nvim-tree/nvim-web-devicons",
     'neovim/nvim-lspconfig',          -- LSP configurations
    'hrsh7th/nvim-cmp',              -- Autocompletion plugin
    'hrsh7th/cmp-nvim-lsp',          -- LSP source for nvim-cmp
    'hrsh7th/cmp-buffer',            -- Buffer completions
    'hrsh7th/cmp-path',              -- Path completions
    'hrsh7th/cmp-cmdline',           -- Command line completions
    'L3MON4D3/LuaSnip',	-- Snippets plugin "nvim-tree/nvim-web-devicons",
  },
  {
    "neovim/nvim-lspconfig",
    config = function()
      -- Removed gopls setup here
    end,
  },
  {
    "catppuccin/nvim",
    name = "catppuccin",
    config = function()
      require("catppuccin").setup({
        flavour = "mocha", -- Set to your preferred flavour
      })
      vim.cmd.colorscheme("catppuccin")
    end,
  },
  {
    "nvim-tree/nvim-tree.lua",
    config = function()
      require("nvim-tree").setup {}
    end,
  },
  {
    "nvim-lualine/lualine.nvim",
    requires = { "nvim-tree/nvim-web-devicons", opt = true },
    config = function()
      require("lualine").setup {
        options = {
          theme = "catppuccin",
          section_separators = { "", "" },
          component_separators = { "", "" },
        },
      }
    end,
  },
  {
    'romgrk/barbar.nvim',
    requires = { 'nvim-tree/nvim-web-devicons' }
  },
  {
    "hrsh7th/nvim-cmp",
    requires = {
      "hrsh7th/cmp-nvim-lsp",    -- LSP source for nvim-cmp
      "hrsh7th/cmp-buffer",      -- Buffer completions
      "hrsh7th/cmp-path",        -- Path completions
      "hrsh7th/cmp-cmdline",     -- Command-line completions
      "L3MON4D3/LuaSnip",        -- Snippet engine
      "saadparwaiz1/cmp_luasnip" -- Snippet completions
    }
  },
})

