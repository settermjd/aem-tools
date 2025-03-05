# AEM Tools

This is a small project, which may or may not grow beyond the initial few commits, that contains tools to help me (and perhaps you) when working with [AEM (Adobe Experience Manager)][aem].

Currently, there's only one tool. 
It converts live blog links into AEM edit links (links used to edit that same post within your AEM instance).

## Getting Started

To use the project, you can either clone it to your local machine and build it, or download the binary for your platform.

If you want to clone and build it, run the commands below, wherever you store your Go projects.

```bash
git clone git@github.com:settermjd/aem-tools.git
cd aem-tools
go install
```

The generated binary will now be available in the bin directory in your Go path, and named _gem-aem-edit-links_.

You can double-check that the file is there by running one of the commands in the example below, assuming that you're using Linux or macOS.

```bash
# Use which to find the binary
which gen-aem-edit-links

// Use ls to print out the file information
ls $(go env GOPATH)/bin/gen-aem-edit-links
```

Regardless of whether you clone and install the binary or download the pre-generated binary, you need to create two files — in the directory where you will run the command:

- A file to store the links that you need to convert to AEM edit links
  This can be named whatever you want.
- A _.env_ file to store the binary's configuration options

In _.env_, paste the following configuration:

```ini
AEM_LINK_TEMPLATE=https://author-<AUTHOR_PREFIX>.adobeaemcloud.com/ui#/aem/editor.html/content/<CONTENT_PATH>/blog/%s.html?appId=aemshell
LINKS_FILE=links.txt
```

`LINKS_FILE` stores the name of the file in the current directory that contain the links that you need to convert.
`AEM_LINK_TEMPLATE` is the AEM URL that you will use to edit the converted link. 
In the example above, you'll need to replace `<AUTHOR_PREFIX>` with your author prefix and `<CONTENT_PATH>` with the URL path up until the placeholder that will be replace with the URL's slug (`%s`).

> [!NOTE]
> See your AEM administrator(s) for these details.

> [!NOTE]
> The AEM link template is a little inflexible at the moment, as it must include `/blog/` before the slug placeholder. 
> That may well change in future versions.
> It's set that way as that's what I'm working with.

[aem]: https://business.adobe.com/products/experience-manager/adobe-experience-manager.html