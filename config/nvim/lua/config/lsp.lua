

local lspconfig = require('lspconfig')

lspconfig.gopls.setup{
    on_attach = function(client, bufnr)
        -- Enable completion triggered by <c-space>
        local opts = { noremap=true, silent=true }
        vim.api.nvim_buf_set_keymap(bufnr, 'n', '<C-Space>', '<Cmd>lua vim.lsp.buf.completion()<CR>', opts)

        -- Other key mappings for LSP
        vim.api.nvim_buf_set_keymap(bufnr, 'n', 'gd', '<Cmd>lua vim.lsp.buf.definition()<CR>', opts)
        vim.api.nvim_buf_set_keymap(bufnr, 'n', 'K', '<Cmd>lua vim.lsp.buf.hover()<CR>', opts)
        vim.api.nvim_buf_set_keymap(bufnr, 'n', '<leader>rn', '<Cmd>lua vim.lsp.buf.rename()<CR>', opts)
        -- Add more key mappings as needed
    end,
    flags = {
        debounce_text_changes = 150,
},
}


local cmp = require('cmp')

vim.lsp.handlers["textDocument/hover"] = vim.lsp.with(vim.lsp.handlers.hover, {
    border = "rounded",  -- Can be 'rounded', 'double', 'solid', etc.
    max_width = 80,      -- Set max width of the hover window
    max_height = 15,     -- Set max height of the hover window
    winhighlight = 'NormalFloat:NormalFloat,FloatBorder:FloatBorder',
})

cmp.setup({
    snippet = {
        expand = function(args)
            require('luasnip').lsp_expand(args.body) -- For snippet support
        end,
    },
    mapping = {
        ['<C-n>'] = cmp.mapping.select_next_item(),
        ['<C-p>'] = cmp.mapping.select_prev_item(),
        ['<C-d>'] = cmp.mapping.scroll_docs(-4),
        ['<C-f>'] = cmp.mapping.scroll_docs(4),
        ['<C-Space>'] = cmp.mapping.complete(),
        ['<C-e>'] = cmp.mapping.close(),
        ['<CR>'] = cmp.mapping.confirm({ select = true }),
    },
    sources = {
        { name = 'nvim_lsp' },
        { name = 'buffer' },
        { name = 'path' },
    },
    window = {
        completion = {
            border = 'rounded',  -- Use 'rounded', 'double', 'solid', etc.
            winhighlight = 'Normal:Pmenu,FloatBorder:Pmenu,CursorLine:CursorLine,Search:None',
            col_offset = 0,
            side_padding = 0,
        },
        documentation = {
            border = 'rounded',
            winhighlight = 'NormalFloat:NormalFloat,FloatBorder:FloatBorder',
        },
    },
})

vim.diagnostic.config({
    virtual_text = {
        prefix = '●',  -- Could be '●', 'x', '!' etc.
        spacing = 4,
    },
    signs = true,
    update_in_insert = true, -- Enable updates in insert mode
    underline = true,
    severity_sort = true,
})
