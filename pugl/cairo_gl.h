/*
  Copyright 2016 David Robillard <http://drobilla.net>

  Permission to use, copy, modify, and/or distribute this software for any
  purpose with or without fee is hereby granted, provided that the above
  copyright notice and this permission notice appear in all copies.

  THIS SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
*/

#ifndef CAIRO_GL_H
#define CAIRO_GL_H

#if defined(PUGL_HAVE_GL) && defined(PUGL_HAVE_CAIRO)

#include <cairo/cairo.h>
#include <stdint.h>

#include "gl.h"

typedef struct {
	unsigned texture_id;
	uint8_t* buffer;
} PuglCairoGL;

cairo_surface_t*
pugl_cairo_gl_create(PuglCairoGL* ctx, int width, int height, int bpp);

void
pugl_cairo_gl_free(PuglCairoGL* ctx);

void
pugl_cairo_gl_configure(PuglCairoGL* ctx, int width, int height);

void
pugl_cairo_gl_draw(PuglCairoGL* ctx, int width, int height);

#endif

#endif // CAIRO_GL_H