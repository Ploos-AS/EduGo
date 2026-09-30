PANDOC ?= pandoc
DIST := dist
NO_SRC := $(shell find course/no -maxdepth 1 -type f -name '*.md' | sort -V)
EN_SRC := $(shell find course/en -maxdepth 1 -type f -name '*.md' | sort -V)

.PHONY: web epub-nb epub-en pdf-nb pdf-en clean

clean:
	rm -rf $(DIST)

web:
	mkdir -p $(DIST)/web/no $(DIST)/web/en
	$(PANDOC) --standalone --toc -o $(DIST)/web/no/index.html $(NO_SRC)
	$(PANDOC) --standalone --toc -o $(DIST)/web/en/index.html $(EN_SRC)

epub-nb:
	mkdir -p $(DIST)
	$(PANDOC) --toc --metadata title=EduGo --metadata author="Per Gustav Ousdal" -o $(DIST)/EduGo-nb.epub $(NO_SRC)

epub-en:
	mkdir -p $(DIST)
	$(PANDOC) --toc --metadata title=EduGo --metadata author="Per Gustav Ousdal" -o $(DIST)/EduGo-en.epub $(EN_SRC)

pdf-nb:
	mkdir -p $(DIST)
	$(PANDOC) --pdf-engine=xelatex --toc --metadata title=EduGo --metadata author="Per Gustav Ousdal" -o $(DIST)/EduGo-nb.pdf $(NO_SRC)

pdf-en:
	mkdir -p $(DIST)
	$(PANDOC) --pdf-engine=xelatex --toc --metadata title=EduGo --metadata author="Per Gustav Ousdal" -o $(DIST)/EduGo-en.pdf $(EN_SRC)
