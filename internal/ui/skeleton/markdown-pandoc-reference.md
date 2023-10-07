---
title: Especificación de Markdown para Pandoc
description: Este es un resumen de la especificación de Markdown para el conversor de formatos Pandoc
date: 06 Octubre 2023
tags: ["markdown", "cheat sheet", "demo"]
---

## Introducción

Pandoc es una biblioteca Haskell para convertir de un formato de marcado a otro, y una herramienta de línea de comandos
que utiliza esta biblioteca.

Pandoc puede convertir entre numerosos formatos de marcado y de procesamiento de textos, incluidos, entre otros, varios
tipos de Markdown, HTML, LaTeX y Word docx. Para ver la lista completa de formatos de entrada y salida, consulte las
opciones `--from` y `--to`. Pandoc también puede generar archivos PDF.

La versión mejorada de Markdown de Pandoc incluye sintaxis para tablas, listas de definiciones, bloques de metadatos,
notas a pie de página, citas, matemáticas y mucho más.

Pandoc tiene un diseño modular: consta de un conjunto de lectores, que analizan el texto en un formato determinado y
producen una representación nativa del documento (un árbol de sintaxis abstracta o AST), y un conjunto de escritores,
que convierten esta representación nativa en un formato de destino. Por tanto, para añadir un formato de entrada o de
salida basta con añadir un lector o un escritor. Los usuarios también pueden ejecutar filtros pandoc personalizados para
modificar el AST intermedio.

Dado que la representación intermedia de pandoc de un documento es menos expresiva que muchos de los formatos entre los
que convierte, no hay que esperar conversiones perfectas entre todos los formatos. Pandoc intenta preservar los
elementos estructurales de un documento, pero no detalles de formato como el tamaño de los márgenes. Y algunos elementos
del documento, como las tablas complejas, pueden no encajar en el modelo de documento simple de pandoc. Aunque las
conversiones del Markdown de pandoc a todos los formatos aspiran a ser perfectas, cabe esperar que las conversiones
desde formatos más expresivos que el Markdown de pandoc tengan pérdidas.

Pandoc entiende una versión ampliada y ligeramente revisada de la sintaxis
[Markdown](https://daringfireball.net/projects/markdown/) de John Gruber. Este documento explica la sintaxis, señalando
las diferencias con el Markdown original.

## Filosofía

Markdown se ha diseñado para que sea fácil de escribir y, lo que es aún más importante, fácil de leer:

> Un documento con formato Markdown debería poder publicarse tal cual, como texto plano, sin que parezca que ha sido
> marcado con etiquetas o instrucciones de formato. -- [John Gruber](https://daringfireball.net/projects/markdown/syntax#philosophy)

Este principio ha guiado las decisiones de pandoc a la hora de encontrar sintaxis para tablas, notas a pie de página y
otras extensiones.

Sin embargo, hay un aspecto en el que los objetivos de pandoc difieren de los objetivos originales de Markdown. Mientras
que Markdown se diseñó originalmente pensando en la generación de HTML, pandoc está pensado para múltiples formatos de
salida. Así, aunque pandoc permite incrustar HTML sin formato, lo desaconseja y ofrece otras formas, no HTML, de
representar elementos importantes del documento, como listas de definiciones, tablas, matemáticas y notas a pie de
página.

## Elementos de markdown

### Párrafos

Un párrafo es una o más líneas de texto seguidas de una o más líneas en blanco. Las nuevas líneas se tratan como
espacios, por lo que puedes redistribuir los párrafos como quieras. Si necesita un salto de línea, coloque dos o más
espacios al final de la línea.

#### Escape de saltos de línea

Una barra invertida (`\\`) seguida de una nueva línea también es un salto de línea.

> **Nota:** en las celdas de tablas multilínea y de cuadrícula, ésta es la única forma de crear un salto de línea duro,
> ya que los espacios finales de las celdas se ignoran.

### Encabezamientos

Existen dos tipos de encabezamientos: Setext y ATX.

#### Encabezamientos estilo Setext

Un encabezamiento setext-style es una línea de texto "subrayada" con una fila de signos `=` (para un encabezamiento de
nivel uno) o signos `-` (para un encabezamiento de nivel dos):

```
Cabecera nivel uno
==================

Cabecera nivel dos
------------------
```

El texto del encabezamiento puede contener formato en línea, como el énfasis (véase *Formato en línea*, más abajo).

#### Encabezamientos de estilo ATX

Un encabezamiento de estilo ATX consta de uno a seis signos `#` y una línea de texto, seguida opcionalmente
de cualquier número de signos `#`. El número de signos `#` al principio de la línea es el nivel
del encabezamiento:

```
## Un encabezamiento de nivel dos

### Un encabezamiento de nivel tres ###
```

Al igual que con los encabezados de estilo setext, el texto del encabezado puede contener formato:

```
## Un encabezamiento de nivel uno con [enlace](/url) y *énfasis
```

#### espacio en blanco antes de la cabecera

La sintaxis original de Markdown no requiere una línea en blanco antes de un título. Markdown Pandoc sí lo exige
(excepto, por supuesto, al principio del documento). La razón de este requisito es que es muy fácil que un `#` termine
al principio de una línea por accidente (quizás por el ajuste de línea). Por ejemplo:

```
Me gustan varios de sus sabores de helado:
#22, por ejemplo, y #5.
```

#### Atributos de cabecera

Esta sintaxis permite asignar atributos a los títulos al final de la línea que contiene el texto del título:

```
{#identifier .class .class key=value key=value}
```

Así, por ejemplo, a los siguientes encabezados se les asignará el identificador [`foo`]{.nowrap}:

```
# Mi encabezamiento {#foo}

## Mi encabezamiento ## {#foo}

Mi otro encabezamiento {#foo}
----------------------
```

*Esta sintaxis es compatible con [PHP Markdown Extra](https://michelf.ca/projects/php-markdown/extra/).*

Tenga en cuenta que, aunque esta sintaxis permite asignar clases y atributos clave/valor, los escritores no suelen
utilizar toda esta información. Los identificadores, las clases y los atributos clave/valor se utilizan en HTML y en
formatos basados en HTML como EPUB y slidy.

Los encabezados con la clase `unnumbered` no se numerarán. Un guión simple (`-`) en un contexto de atributo equivale a
[`.unnumbered`], y es preferible en documentos que no estén en inglés. Así,

```
# Mi encabezamiento {-}
```

es lo mismo que

```
# Mi encabezamiento {.unnumbered}
```

Si la clase `unlisted` está presente además de `unnumbered`, el encabezamiento no se incluirá en una tabla de contenidos.

#### Referencias implícitas de cabecera

Pandoc Markdown se comporta como si se hubieran definido enlaces de referencia para cada encabezamiento. Así, para
enlazar con un encabezamiento

```
# Heading identifiers in HTML
```

puede escribir simplemente

```
[Heading identifiers in HTML]
```

o

```
[Heading identifiers in HTML][]
```

o

```
[the section on heading identifiers][heading identifiers in HTML]
```

en lugar de dar el identificador explícitamente:

```
[Heading identifiers in HTML](#heading-identifiers-in-html)
```

Si hay varios títulos con el mismo texto, la referencia correspondiente sólo enlazará con el primero, y tendrá que
utilizar enlaces explícitos para enlazar con los demás, como se ha descrito anteriormente.

Al igual que los enlaces de referencia normales, estas referencias no distinguen entre mayúsculas y minúsculas.

Las definiciones de referencias explícitas siempre tienen prioridad sobre las referencias implícitas. Así, en el
siguiente ejemplo, el enlace apuntará a `bar`, no a `#foo`:

```
# Foo

[foo]: bar

See [foo]
```

### Citas en bloque

Markdown utiliza convenciones de correo electrónico para citar bloques de texto. Una cita en bloque consiste en uno o
más párrafos u otros elementos en bloque (como listas o títulos), con cada línea precedida por un carácter `>` y un
espacio opcional. (No es necesario que el carácter `>` empiece en el margen izquierdo, pero no debe sangrar más de tres
espacios).

```
> Se trata de una cita en bloque. Este párrafo
> tiene dos líneas.
>
> 1. Esto es una lista dentro de una cita en bloque.
> 2. Segundo punto.
```

También se permite una forma "perezosa", que requiere el carácter `>` sólo en la primera línea de cada bloque:

```
> Esto es una cita de bloque. Este párrafo
tiene dos líneas.
> 1. Esto es una lista dentro de una cita en bloque.
2. Segundo elemento.
```

Entre los elementos de bloque que puede contener una cita de bloque se encuentran otras citas de bloque. Es decir, las
comillas de bloque se pueden anidar:

```
> Esto es una cita en bloque.
>
> > Una cita dentro de otra cita.
```

Si el carácter `>` va seguido de un espacio opcional, ese espacio se considerará parte del marcador de comillas de
bloque y no parte de la sangría del contenido. Así, para poner un bloque de código con sangría en una cita de bloque, se
necesitan cinco espacios después del `>`:

```
>_____código
```

*(Asumiendo el caracter `_` como un espacio)*

##### Espacio en blanco antes del bloque de cita

La sintaxis original de Markdown no requiere una línea en blanco antes de una cita en bloque. Pandoc sí lo exige
(excepto, por supuesto, al principio del documento). La razón de este requisito es que es muy fácil que un `>` termine
al principio de una línea por accidente (quizás por el ajuste de línea). lo siguiente no produce una cita de bloque
anidada:

```
> Esto es una cita de bloque.
>> Esta cita no está anidada.
```


### Bloques Verbatim (código fuente)

#### Bloques de código indentado

Un bloque de texto indentado a cuatro espacios (o una tabulación) se trata como texto literal: es decir, los caracteres
especiales no activan el formato especial, y se conservan todos los espacios y saltos de línea. Por ejemplo

```go
    if (a > 3) {
      moverBarco(5 * gravedad, ABAJO);
    }
```

La sangría inicial (cuatro espacios o un tabulador) no se considera parte del texto literal, y se elimina en la salida.

> **Nota:** las líneas en blanco en el texto literal no necesitan comenzar con cuatro espacios.

#### Bloques de código delimitados

Además de los bloques de código con sangría estándar, pandoc admite bloques de código *delimitados*. Éstos comienzan con
una fila de tres o más tildes (se conoce también como circunflejo) (`~`) y terminan con una fila de tildes que debe ser
al menos tan larga como la fila inicial. Todo lo que hay entre estas líneas se trata como código. No es necesario
indentar:

    ~~~~~~~
    if (a > 3) {
      moveShip(5 * gravity, DOWN);
    }
    ~~~~~~~

Al igual que los bloques de código normales, los bloques de código cercados deben separarse del texto circundante
mediante líneas en blanco.

Si el propio código contiene una fila de tildes o barras invertidas, basta con utilizar una fila más larga de tildes o
barras invertidas al principio y al final:

    ~~~~~~~~~~~~~~~~
    ~~~~~~~~~~
    código con tildes
    ~~~~~~~~~~
    ~~~~~~~~~~~~~~~~

También se puede utilizar backticks (`` ` ``) en lugar de tildes (`~`)

    ```go
    import "websocket"

    type T struct {
        Msg string
        Count int
    }

    // receive JSON type T
    var data T
    websocket.JSON.Receive(ws, &data)

    // send JSON type T
    websocket.JSON.Send(ws, data)
    ```

Opcionalmente, puede adjuntar atributos al bloque de código cercado o backtick utilizando esta sintaxis:

    ~~~~ {#mycode .haskell .numberLines startFrom="100"}
    qsort []     = []
    qsort (x:xs) = qsort (filter (< x) xs) ++ [x] ++
                   qsort (filter (>= x) xs)
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Aquí `mycode` es un identificador, `haskell` y `numberLines` son clases, y `startFrom` es un atributo con valor 100.
Algunos formatos de salida pueden utilizar esta información para resaltar la sintaxis. De lo contrario, el bloque de
código anterior aparecerá como sigue:

```html
<pre id="mycode" class="haskell numberLines" startFrom="100">
  <code>
  ...
  </code>
</pre>
```

La clase `numberLines` (o `number-lines`) hará que las líneas del bloque de código se numeren, empezando por 1 o el
valor del atributo `startFrom`. La clase `lineAnchors` (o `line-anchors`) hará que las líneas sean anclas clickables en
la salida HTML.

También se puede utilizar una forma abreviada para especificar el idioma del bloque de código:

    ```haskell
    qsort [] = []
    ```

Esto es equivalente a:

    ``` {.haskell}
    qsort [] = []
    ```

Esta forma abreviada puede combinarse con atributos:

    ```haskell {.numberLines}
    qsort [] = []
    ```

Que es equivalente a:

    ``` {.haskell .numberLines}
    qsort [] = []
    ```


### Líneas de bloque

Un bloque de líneas es una secuencia de líneas que comienza con una barra vertical (`|`) seguida de un espacio. La
división en líneas se conservará en la salida, al igual que los espacios iniciales; en caso contrario, las líneas se
formatearán como Markdown. Esto es útil para versos y direcciones:

```
| El limerick empaqueta risas anatómicas
| En el espacio que es bastante económico.
| Pero los buenos que he visto...
| Rara vez son limpios
| Y los limpios rara vez son cómicos.

| 200 Main St.
| Berkeley, CA 94718
```

Si es necesario, las líneas pueden ir seguidas, pero la línea de continuación debe comenzar con un espacio.

```
| El Muy Honorable Muy Venerable y Recto Samuel L.
 Constable, Jr.
| 200 Main St.
| Berkeley, CA 94718
```

El formato en línea (como el énfasis) está permitido en el contenido, pero no el formato a nivel de bloque (como las
citas en bloque o las listas).

### listas

#### Listas con viñetas

Una lista con viñetas es una lista de elementos con viñetas. Un elemento de una lista con viñetas comienza con una
viñeta (*, + o -). He aquí un ejemplo sencillo:

```
* uno
* dos
* tres
```

Así obtendrá una lista "compacta". Si desea una lista "suelta", en la que cada elemento tenga formato de párrafo, ponga
espacios entre los elementos:

```
* uno

* dos

* tres
```

No es necesario que las viñetas estén alineadas con el margen izquierdo; pueden tener una sangría de uno, dos o tres
espacios. La viñeta debe ir seguida de un espacio en blanco.

Los elementos de la lista se ven mejor si las líneas siguientes están a ras de la primera línea (después de la viñeta):

```
* Aquí está mi primer
  elemento de la lista.
* y el segundo.
```

Pero Markdown también permite un formato "perezoso":

```
* aquí está mi primer
elemento de la lista.
* y mi segundo.
```

#### Contenido en bloque en los elementos de lista

Un elemento de lista puede contener varios párrafos y otros contenidos en bloque. Sin embargo, los párrafos siguientes
deben ir precedidos de una línea en blanco y sangrados para alinearse con el primer contenido sin espacios después del
marcador de lista.

```
* Primer párrafo.

  Continúa.

* Segundo párrafo. Con un bloque de código, que debe ser sangrado
  ocho espacios:

{ code }
```
Excepción: si el marcador de lista va seguido de un bloque de código sangrado, que debe comenzar 5 espacios después del
marcador de lista, los párrafos siguientes deben comenzar dos columnas después del último carácter del marcador de
lista:

```
* code

párrafo de continuación
```

Los elementos de lista pueden incluir otras listas. En este caso, la línea en blanco anterior es opcional. La lista
anidada debe estar sangrada para alinearse con el primer carácter sin espacio después del marcador de lista del elemento
de lista que la contiene.

```
* Frutas
  + manzanas
    - macintosh
    - rojo delicioso
  + peras
  + melocotones
* verduras
  + brócoli
  + acelgas
```

Como ya se ha indicado, Markdown permite escribir elementos de lista "perezosamente", en lugar de sangrar las líneas de
continuación. Sin embargo, si hay varios párrafos u otros bloques en un elemento de lista, la primera línea de cada uno
debe sangrarse.

```
+ Un elemento de lista
perezoso.

+ Otro; este se ve
mal pero es legal.

    Segundo párrafo del segundo
punto de la lista.
```


#### Listas ordenadas

Las listas ordenadas funcionan igual que las listas con viñetas, salvo que los elementos comienzan con enumeradores en
lugar de con viñetas.

En el Markdown original, los enumeradores son números decimales seguidos de un punto y un espacio. Los números en sí se
ignoran, por lo que no hay diferencia entre esta lista:

```
1. uno
2. dos
3. tres
```

y ésta:

```
5. uno
7. dos
1. tres
```

A diferencia del Markdown original, el Markdown de Pandoc permite marcar los elementos ordenados de una lista con letras
mayúsculas y minúsculas y números romanos, además de números arábigos. Los marcadores de lista pueden ir encerrados
entre paréntesis o seguidos de un único paréntesis derecho o punto. Deben estar separados del texto que les sigue por al
menos un espacio y, si el marcador de lista es una letra mayúscula con un punto, por al menos dos espacios. El objetivo
de esta regla es garantizar que los párrafos normales que comienzan con las iniciales de las personas, como

```
B. Russell ganó un Premio Nobel.
```

no se traten como elementos de lista.

Esta regla no impedirá que

```
(C) 2007 Joe Smith
```

se interprete como un elemento de lista. En este caso, puede utilizarse una barra invertida:

```
(C\) 2007 Joe Smith
```

También permite utilizar "#" como marcador de lista ordenada en lugar de un número:

```
#. uno
#. dos
#. tres
#. etc.
```

El Markdown de Pandoc también presta atención al tipo de marcador de lista utilizado y al número inicial, y ambos se
conservan siempre que es posible en el formato de salida. Así, se obtiene una lista con números seguidos de un
paréntesis, empezando por 9, y una sublista con números romanos en minúsculas:

```
9) Noveno
10) Décimo
11) Undécimo
i. subuno
ii. subdos
iii. subtres
```

Pandoc iniciará una nueva lista cada vez que se utilice un tipo diferente de marcador de lista. Así, lo siguiente creará
tres listas:

```
(2) Dos
(5) Tres
1. Cuatro
* Cinco
```

#### Listas de tareas

Pandoc soporta listas de tareas, utilizando la sintaxis de [GitHub-Flavored
Markdown](https://github.github.com/gfm/#task-list-items-extension-).

```
- [ ] un elemento no marcado de la lista de tareas
- [x] elemento marcado
```

#### Listas de definición

Pandoc admite listas de definiciones, utilizando la sintaxis de PHP Markdown Extra con algunas extensiones

```
Término 1

:   Definición 1

Término 2 con *marcado en línea*

:   Definición 2

        { algún código, parte de la Definición 2 }

    Tercer párrafo de la definición 2.
```

Cada término debe caber en una línea, que puede ir seguida opcionalmente de una línea en blanco, y debe ir seguida de
una o más definiciones. Una definición comienza con dos puntos o tilde, que pueden ir sangrados uno o dos espacios.

Un término puede tener varias definiciones, y cada definición puede constar de uno o varios elementos de bloque
(párrafo, bloque de código, lista, etc.), cada uno de ellos sangrado cuatro espacios o un tabulador. El cuerpo de la
definición (sin incluir la primera línea) debe tener una sangría de cuatro espacios. Sin embargo, al igual que ocurre
con otras listas de Markdown, puede omitirse la sangría, excepto al principio de un párrafo u otro elemento de bloque:

```
Término 1

:   Definición
con continuación perezosa.

    Segundo párrafo de la definición.
```

Si deja espacio antes de la definición (como en el ejemplo anterior), el texto de la definición se tratará como un
párrafo. En algunos formatos de salida, esto significará un mayor espaciado entre los pares término/definición. Para
obtener una lista de definiciones más compacta, omita el espacio antes de la definición:

```
Término 1
 ~ Definición 1

Término 2
 ~ Definición 2a
 ~ Definición 2b
```

Tenga en cuenta que se requiere espacio entre los elementos de una lista de definiciones.

#### Listas de ejemplos

El marcador especial de lista `@` puede utilizarse para ejemplos numerados secuencialmente. El primer elemento de la
lista con un marcador @ se numerará como "1", el siguiente como "2", y así sucesivamente en todo el documento. No es
necesario que los ejemplos numerados aparezcan en una sola lista; cada nueva lista que utilice `@` continuará donde se
detuvo la anterior. Así, por ejemplo:

```
(@) Mi primer ejemplo se numerará (1).
(@) Mi segundo ejemplo se numerará (2).

Explicación de los ejemplos....

(@) Mi tercer ejemplo irá numerado (3).
```

Los ejemplos numerados pueden etiquetarse y citarse en otras partes del documento:

```
(@bueno) Éste es un buen ejemplo.

Como ilustra (@bueno), ...
```

La etiqueta puede ser cualquier cadena de caracteres alfanuméricos, guiones bajos o guiones.

> **Nota:** los párrafos de continuación de las listas de ejemplo siempre deben tener una sangría de cuatro espacios,
> independientemente de la longitud del marcador de lista.


#### Finalizar una lista

¿Y si quieres poner un bloque de código sangrado después de una lista?

```
- elemento uno
- elemento dos

{ mi bloque de código }
```

¡Problema! Aquí pandoc (como otras implementaciones de Markdown) tratará `{ mi bloque de código }` como el segundo párrafo
del elemento dos, y no como un bloque de código.

Para "cortar" la lista después del punto dos, puede insertar algún contenido no sangrado, como un comentario HTML, que
no producirá una salida visible en ningún formato:

```
- elemento uno
- elemento dos

<!-- fin de la lista -->

{ mi bloque de código }
```

Puedes usar el mismo truco si quieres dos listas consecutivas en lugar de una lista grande:

```
1. una
2. dos
3. tres

<!-- -->

1. uno
2. dos
3. tres
```

### Líneas horizontales

Una línea que contenga una fila de tres o más caracteres *, - o _ (opcionalmente separados por espacios) produce una
regla horizontal:

```

* * * *

---------------


```
Recomendamos encarecidamente que las reglas horizontales estén separadas del texto circundante por líneas en blanco. Si
una regla horizontal no va seguida de una línea en blanco, pandoc puede intentar interpretar las líneas que siguen como
un bloque de metadatos YAML o una tabla.


### Tablas

Se pueden utilizar cuatro tipos de tablas. En los cuatro tipos de tablas (como se ilustra en los ejemplos siguientes)
puede incluirse opcionalmente un pie de ilustración. Una leyenda es un párrafo que comienza con la cadena Table: (o
table: o simplemente :), que se eliminará. Puede aparecer antes o después de la tabla.

#### Tabla simple

Las tablas sencillas tienen este aspecto:

```
  Right     Left     Center     Default
-------     ------ ----------   -------
     12     12        12            12
    123     123       123          123
      1     1          1             1

Table:  Demonstration of simple table syntax.
```

La cabecera y las filas de la tabla deben caber cada una en una línea. La alineación de las columnas viene determinada
por la posición del texto de cabecera con respecto a la línea discontinua situada debajo:3

* Si la línea discontinua está alineada con el texto de cabecera a la derecha, pero lo sobrepasa a la izquierda, la
  columna está alineada a la derecha.
* Si la línea discontinua está alineada con el texto de cabecera a la izquierda, pero lo sobrepasa a la derecha, la
  columna está alineada a la izquierda.
* Si la línea discontinua sobrepasa el texto de cabecera por ambos lados, la columna está centrada.
* Si la línea discontinua está alineada con el texto de la cabecera en ambos lados, se utiliza la alineación por defecto
  (en la mayoría de los casos, será la izquierda).
* La tabla debe terminar con una línea en blanco o una línea de guiones seguida de una línea en blanco.

Puede omitirse la fila de encabezado de columna, siempre que se utilice una línea de guiones para finalizar la tabla.
Por ejemplo:

```
-------     ------ ----------   -------
     12     12        12             12
    123     123       123           123
      1     1          1              1
-------     ------ ----------   -------
```

Cuando se omite la fila de cabecera, la alineación de las columnas se determina en función de la primera línea del
cuerpo de la tabla. Así, en las tablas anteriores, las columnas estarían alineadas a la derecha, izquierda, centro y
derecha, respectivamente.

#### Tabla multilínea

Las tablas multilínea permiten que el encabezado y las filas de la tabla abarquen varias líneas de texto (pero no
admiten celdas que abarquen varias columnas o filas de la tabla). He aquí un ejemplo:

```
-------------------------------------------------------------
 Centered   Default           Right Left
  Header    Aligned         Aligned Aligned
----------- ------- --------------- -------------------------
   First    row                12.0 Example of a row that
                                    spans multiple lines.

  Second    row                 5.0 Here's another one. Note
                                    the blank line between
                                    rows.
-------------------------------------------------------------

Table: Here's the caption. It, too, may span
multiple lines.
```

Funcionan como las tablas simples, pero con las siguientes diferencias:

* Deben comenzar con una fila de guiones, antes del texto de cabecera (a menos que se omita la fila de cabecera).
* Deben terminar con una fila de guiones y, a continuación, una línea en blanco.
* Las filas deben estar separadas por líneas en blanco.
* En las tablas multilínea, el analizador de tablas presta atención a las anchuras de las columnas, y los escritores
  intentan reproducir estas anchuras relativas en la salida. Por lo tanto, si encuentra que una de las columnas es
  demasiado estrecha en la salida, intente ensancharla en la fuente Markdown.

La cabecera puede omitirse tanto en las tablas multilínea como en las tablas simples:

```
----------- ------- --------------- -------------------------
   First    row                12.0 Example of a row that
                                    spans multiple lines.

  Second    row                 5.0 Here's another one. Note
                                    the blank line between
                                    rows.
----------- ------- --------------- -------------------------

: Here's a multiline table without a header.
```

Es posible que una tabla multilínea sólo tenga una fila, pero la fila debe ir seguida de una línea en blanco (y luego la
fila de guiones que finaliza la tabla), o la tabla puede interpretarse como una tabla simple.

#### Tabla cuadrícula

Las tablas de cuadrícula tienen este aspecto:

```
: Sample grid table.

+---------------+---------------+--------------------+
| Fruit         | Price         | Advantages         |
+===============+===============+====================+
| Bananas       | $1.34         | - built-in wrapper |
|               |               | - bright color     |
+---------------+---------------+--------------------+
| Oranges       | $2.10         | - cures scurvy     |
|               |               | - tasty            |
+---------------+---------------+--------------------+
```

La fila de *=s* separa la cabecera del cuerpo de la tabla, y puede omitirse para una tabla sin cabecera. Las celdas de
las tablas cuadriculadas pueden contener elementos de bloque arbitrarios (párrafos múltiples, bloques de código, listas,
etc.).

Las celdas pueden abarcar varias columnas o filas:

```
+---------------------+----------+
| Property            | Earth    |
+=============+=======+==========+
|             | min   | -89.2 °C |
| Temperature +-------+----------+
| 1961-1990   | mean  | 14 °C    |
|             +-------+----------+
|             | max   | 56.7 °C  |
+-------------+-------+----------+
```

Una cabecera de tabla puede contener más de una fila:

```
+---------------------+-----------------------+
| Location            | Temperature 1961-1990 |
|                     | in degree Celsius     |
|                     +-------+-------+-------+
|                     | min   | mean  | max   |
+=====================+=======+=======+=======+
| Antarctica          | -89.2 | N/A   | 19.8  |
+---------------------+-------+-------+-------+
| Earth               | -89.2 | 14    | 56.7  |
+---------------------+-------+-------+-------+
```

Las alineaciones se pueden especificar como con las tablas de tuberías, poniendo dos puntos en los límites de la línea
separadora después de la cabecera:

```
+---------------+---------------+--------------------+
| Right         | Left          | Centered           |
+==============:+:==============+:==================:+
| Bananas       | $1.34         | built-in wrapper   |
+---------------+---------------+--------------------+
```

En las tablas sin encabezado, los dos puntos van en la línea superior:

```
+--------------:+:--------------+:------------------:+
| Right         | Left          | Centered           |
+---------------+---------------+--------------------+
```

Un pie de tabla puede definirse encerrándolo con líneas separadoras que utilicen `=` en lugar de `-`:

```
 +---------------+---------------+
 | Fruit         | Price         |
 +===============+===============+
 | Bananas       | $1.34         |
 +---------------+---------------+
 | Oranges       | $2.10         |
 +===============+===============+
 | Sum           | $3.44         |
 +===============+===============+
```

El pie debe colocarse siempre en la parte inferior de la tabla.

#### Tabla de tuberías (pipes)

Las tablas de tuberías tienen este aspecto:

```
| Right | Left | Default | Center |
|------:|:-----|---------|:------:|
|   12  |  12  |    12   |    12  |
|  123  |  123 |   123   |   123  |
|    1  |    1 |     1   |     1  |

: Demonstration of pipe table syntax.
```

La sintaxis es idéntica a la de las [tablas PHP Markdown Extra](https://michelf.ca/projects/php-markdown/extra/#table).
Los caracteres de tubería al principio y al final son opcionales, pero se requieren tuberías entre todas las columnas.
Los dos puntos indican la alineación de las columnas. La cabecera no puede omitirse. Para simular una tabla sin
encabezado, incluya un encabezado con celdas en blanco.

Como los tubos indican los límites de las columnas, no es necesario alinearlas verticalmente, como en el ejemplo
anterior. Por lo tanto, se trata de una tabla de tubos perfectamente legal (aunque fea):

```
fruit| price
-----|-----:
apple|2.05
pear|1.37
orange|3.09
```

Las celdas de las tablas pipe no pueden contener elementos de bloque como párrafos y listas, y no pueden abarcar varias
líneas. Si alguna línea de la fuente de markdown es más larga que el ancho de columna, entonces la tabla ocupará todo el
ancho del texto y el contenido de las celdas se envolverá, con el ancho relativo de las celdas determinado por el número
de guiones en la línea que separa la cabecera de la tabla del cuerpo de la tabla. (Por ejemplo ---|- haría que la
primera columna ocupara 3/4 y la segunda 1/4 del ancho total del texto). Por otro lado, si ninguna línea es más ancha
que el ancho de la columna, el contenido de las celdas no se envolverá, y las celdas tendrán el tamaño de su contenido.
Pandoc también reconoce tablas de tuberías de la siguiente forma:

```
| One | Two   |
|-----+-------|
| my  | table |
| is  | nice  |
```

La diferencia es que se utiliza `+` en lugar de `|`. Otras características de orgtbl no están soportadas. En particular,
para obtener una alineación de columnas no predeterminada, tendrá que añadir dos puntos como se indica más arriba.


### Bloque de metadatos

Un bloque de metadatos [YAML](https://yaml.org/spec/1.2/spec.html) es un objeto YAML válido, delimitado por una línea de
tres guiones (`---`) en la parte superior y una línea de tres guiones (`---`) o tres puntos (`...`) en la parte
inferior. La línea inicial `---` no debe ir seguida de una línea en blanco. Un bloque de metadatos YAML puede aparecer
en cualquier parte del documento, pero si no está al principio, debe ir precedido de una línea en blanco.

Los metadatos se tomarán de los campos del objeto YAML y se añadirán a los metadatos existentes en el documento. Los
metadatos pueden contener listas y objetos (anidados arbitrariamente), pero todos los escalares de cadena se
interpretarán como Markdown. Los campos cuyos nombres terminen en guión bajo serán ignorados por pandoc. (Los nombres de
campo no deben poder interpretarse como números YAML o valores booleanos (por ejemplo, `yes`, `True` y `15` no pueden
utilizarse como nombres de campo).

Un documento puede contener varios bloques de metadatos. Si dos bloques de metadatos intentan establecer el mismo campo,
se tomará el valor del segundo bloque.

Cada bloque de metadatos se gestiona internamente como un documento YAML independiente. Esto significa, por ejemplo, que
cualquier ancla YAML definida en un bloque no puede ser referenciada en otro bloque.

Tenga en cuenta que deben respetarse las reglas de escape de YAML. Así, por ejemplo, si un título contiene dos puntos,
debe entrecomillarse, y si contiene una barra invertida de escape, hay que asegurarse de que no se trata como una
secuencia de escape YAML. El carácter pipe (`|`) puede utilizarse para iniciar un bloque sangrado que se interpretará
literalmente, sin necesidad de escape. Esta forma es necesaria cuando el campo contiene líneas en blanco o formato a
nivel de bloque:

```
---
title:  'This is the title: it contains a colon'
author:
- Author One
- Author Two
keywords: [nothing, nothingness]
abstract: |
  This is the abstract.

  It consists of two paragraphs.
...

## This is a subtitle

This is a paragraph.
```

El bloque literal después del `|` debe tener sangría relativa a la línea que contiene el `|`. Si no es así, el YAML no
será válido y pandoc no lo interpretará como metadatos. Para una visión general de las complejas reglas que rigen YAML,
consulte la [entrada de Wikipedia sobre sintaxis YAML](https://es.wikipedia.org/wiki/YAML).

Las variables de plantilla se establecerán automáticamente a partir de los metadatos. Así, por ejemplo, al escribir
HTML, la variable abstract se establecerá al equivalente HTML del Markdown en el campo abstract:

```html
<p>Este es el resumen.</p>
<p>Se compone de dos párrafos.</p>
```

Las variables pueden contener estructuras YAML arbitrarias, pero la plantilla debe coincidir con esta estructura. La
variable author en las plantillas por defecto espera una lista simple o una cadena, pero puede cambiarse para soportar
estructuras más complicadas. La siguiente combinación, por ejemplo, añadiría una afiliación al autor si se da una:

```
---
title: The document title
author:
- name: Author One
  affiliation: University of Somewhere
- name: Author Two
  affiliation: University of Nowhere
...

## This is a subtitle

This is a paragraph.
```

Para utilizar los autores estructurados del ejemplo anterior, necesitaría una plantilla personalizada:

```
$for(author)$
    $if(author.name)$
        $author.name$$if(author.affiliation)$ ($author.affiliation$)$endif$
    $else$
        $author$
    $endif$
$endfor$
```

### Escape de caracteres especiales

Excepto dentro de un bloque de código o código en línea, cualquier carácter de puntuación o espacio precedido de una
barra invertida (`\\`) se tratará literalmente, incluso si normalmente indicaría formato. Así, por ejemplo, si se
escribe:

```
*\*hello\**
```

se obtendrá

```
*\*hello\**
```

en lugar de

```
*\*hello\**
```

Esta regla es más fácil de recordar que la regla original de Markdown, que sólo permite que los siguientes caracteres
sean backslash-escapados:

```
\`*_{}[]()>#+-.!
```

Un espacio con barra invertida se interpreta como un espacio que no se rompe. En la salida TeX, aparecerá como ~. En la
salida HTML y XML, aparecerá como un carácter de espacio sin ruptura unicode literal (tenga en cuenta que, por lo tanto,
parecerá "invisible" en la fuente HTML generada.

Un salto de línea con barra invertida (es decir, una barra invertida al final de una línea) se interpreta como un salto
de línea. Aparecerá en la salida TeX como `\\` y en HTML como `<br />`. Es una buena alternativa a la forma "invisible"
de Markdown de indicar los saltos de línea mediante dos espacios al final de la línea.

Las barras invertidas no funcionan en contextos literales (Bloques de código).

### Formato en línea

#### Énfasis

Para enfatizar algún texto, rodéalo de `*`s o `_`, así:

```
Este texto está _enfatizado con guiones bajos_, y éste
está *enfatizado con asteriscos*.
```

El doble `*` o `_` produce un **énfasis fuerte**:

```
Esto es **énfasis fuerte** y __con guiones bajos__.
```

Un carácter `*` o `_` rodeado de espacios, o con barra invertida, no produce énfasis:

```
Esto es * no acentuado *, y \*tampoco esto\*.
```

Dado que `_` se utiliza a veces dentro de palabras e identificadores, pandoc no interpreta un `_` rodeado de caracteres
alfanuméricos como un marcador de énfasis. Si desea enfatizar sólo parte de una palabra, utilice `*`:

```
feas*ible*, not feas*able*.
```

#### Efecto de tachado

Para tachar una sección de texto con una línea horizontal, comience y termine con `~~`. Así, por ejemplo:

```
This ~~is deleted text~~.
```

#### Superíndices y subíndices

Los superíndices pueden escribirse rodeando el texto en superíndice con caracteres `^`; los subíndices pueden escribirse
rodeando el texto en subíndice con caracteres `~`. Así, por ejemplo,

```
H~2~O es un líquido. 2^10^ es 1024.
```

El texto entre `^...^` o `~...~` no puede contener espacios ni nuevas líneas. Si el texto en superíndice o subíndice
contiene espacios, éstos deben escaparse con barras invertidas. (Esto se hace para evitar superíndices y subíndices
accidentales mediante el uso ordinario de `~` y `^`, y también malas interacciones con las notas a pie de página). Así,
si desea la letra P con "un gato" en los subíndices, utilice `P~a\ gato~`, no `P~a gato~`.

#### Texto literal

Para que un fragmento de texto sea literal, introdúzcalo entre backticks(`` ` ``)

```
¿Cuál es la diferencia entre `>>=` y `>>`?
```

Si el texto literal incluye un backtick, utilice dos backticks:

```
Aquí tienes un backtick literal `` ` ``.
```

(Se ignorarán los espacios después de las comillas de apertura y antes de las comillas de cierre).

La regla general es que un tramo literal empiece con una cadena de puntos y comas consecutivos (opcionalmente seguidos
de un espacio) y termine con una cadena del mismo número de puntos y comas (opcionalmente precedidos de un espacio).

Tenga en cuenta que las barras invertidas (y otras construcciones de Markdown) no funcionan en contextos literales:

```
Se trata de una barra invertida seguida de un asterisco: `\*`.
```

Se pueden adjuntar atributos al texto literal, igual que con los bloques de código vallados:

```
`<$>`{.php}
```

#### Subrayado

Para subrayar texto, utilice la clase de subrayado:

```
[Subrayado]{.underline}
```

Genera el siguiente resultado:

```html
<span class="underline">Subrayado</span>
```

#### Pequeñas capitalizaciones (versalitas)

Para escribir versalitas, utilice la clase `smallcaps`:

```
[Small Caps]{.smallcaps}
```

Genera el siguiente resultado:

```
<span class="smallcaps">Small Caps</span>
```

#### Texto en HTML

Markdown permite insertar HTML sin formato (o DocBook) en cualquier parte de un documento (excepto en contextos
literales, donde <, > y & se interpretan literalmente).

```
This is a <b>bold HTML</b> text and a **bold Markdown** text.
```

#### Enlaces (Links)

Markdown permite especificar los enlaces de varias formas.

##### Enlaces automáticos

Si encierras una URL o una dirección de correo electrónico entre corchetes puntiagudos (`<...>`), se convertirá en un
enlace:

```
<https://google.com>
<sam@green.eggs.ham>
```

##### Enlaces en línea

Un enlace en línea consiste en el texto del enlace entre corchetes cuadrados (`[]`), seguido de la URL entre paréntesis.
(Opcionalmente, la URL puede ir seguida de un título de enlace, entre comillas).

```
Esto es un [enlace en línea](/url), y aquí hay
[uno con un título](https://fsf.org "¡haga clic aquí para pasar un buen rato!").
```

No puede haber ningún espacio entre la parte entre corchetes y la parte entre paréntesis. El texto del enlace puede
contener formato (como énfasis), pero el título no.

Las direcciones de correo electrónico de los enlaces en línea no se detectan automáticamente, por lo que deben ir
precedidas de mailto:

```
[¡Escríbeme!](mailto:sam@green.eggs.ham)
```

##### Enlaces de referencia

Un enlace de referencia explícita consta de dos partes: el enlace propiamente dicho y la definición del enlace, que
puede aparecer en otra parte del documento (antes o después del enlace).

El enlace consta de un texto entre corchetes, seguido de una etiqueta entre corchetes. (No puede haber espacio entre
ambos). La definición del enlace consiste en la etiqueta entre corchetes, seguida de dos puntos y un espacio, seguida de
la URL, y opcionalmente (después de un espacio) un título del enlace entre comillas o entre paréntesis. La etiqueta no
debe ser analizable como una cita (suponiendo que la extensión citations esté activada): las citas tienen prioridad
sobre las etiquetas de los enlaces.

He aquí algunos ejemplos:

```
[my label 1]: /foo/bar.html  "My title, optional"
[my label 2]: /foo
[my label 3]: https://fsf.org (The Free Software Foundation)
[my label 4]: /bar#special  'A title in single quotes'
```

La URL puede ir opcionalmente entre paréntesis angulares:

```
[my label 5]: <http://foo.bar.baz>
```

El título puede ir en la línea siguiente:

```
[my label 3]: https://fsf.org
  "The Free Software Foundation"
```

Tenga en cuenta que las etiquetas de los enlaces no distinguen entre mayúsculas y minúsculas. Por lo tanto, esto
funcionará:

```
Here is [my link][FOO]

[Foo]: /bar/baz
```

En un enlace de referencia implícito, el segundo par de corchetes está vacío:

```
See [my website][].

[my website]: http://foo.bar.baz
```

> *Nota:* Así que lo siguiente está bien en pandoc, aunque no en la mayoría de las otras implementaciones:
>     > My block [quote].
>     >
>     > [quote]: /foo


##### Enlaces internos

Para enlazar con otra sección del mismo documento, utilice el identificador generado automáticamente) o definiéndolo
manualmente. Por ejemplo:

```
## Introducción {#intro}

Véase la [Introducción](#intro).
```

o

```
## Introducción

Véase la [Introducción].

[Introducción]: #introducción
```

##### Imágenes

Un enlace precedido inmediatamente por un `!` será tratado como una imagen. El texto del enlace se utilizará como texto
alternativo de la imagen:

```
![la lune](/path/to/lalune.jpg "Voyage to the moon")

![movie reel]

[movie reel]: /path/to/movie.gif
```

Una imagen con un texto alternativo no vacío, que aparezca sola en un párrafo, se mostrará como una figura con un pie de
foto. El texto alternativo de la imagen se utilizará como pie de foto.

```
![Este es el título](/url/de/imagen.png)
```

Si sólo desea una imagen en línea, asegúrese de que no es lo único que aparece en el párrafo. Una forma de hacerlo es
insertar un espacio de no ruptura después de la imagen:

```
![Esta imagen no será una figura](/url/de/imagen.png)\
```

Se pueden establecer atributos en enlaces e imágenes:

```
An inline ![image](foo.jpg){#id .class width=30 height=20px}
and a reference ![image][ref] with attributes.

[ref]: foo.jpg "optional title" {#id .class key=val key2="val 2"}
```

(Esta sintaxis es compatible con [PHP Markdown Extra](https://michelf.ca/projects/php-markdown/extra/) cuando sólo se
utilizan #id y .class).

Los atributos de anchura y altura de las imágenes reciben un tratamiento especial. Cuando se utilizan sin unidad, se
asume que la unidad es el píxel. Sin embargo, se puede utilizar cualquiera de los siguientes identificadores de unidad:
*px*, *cm*, *mm*, *in*, *inch* y *%*. No debe haber espacios entre el número y la unidad. Por ejemplo:

```
![](archivo.jpg){ width=50% }
```

* Las dimensiones pueden convertirse a una forma compatible con el formato de . La conversión entre píxeles y medidas
  físicas se ve afectada por su dpi (por defecto, se asume 96 dpi, a menos que la propia imagen contenga información
  dpi).
* La unidad % es generalmente relativa a algún espacio disponible. Por ejemplo el ejemplo anterior se renderizará de la
  siguiente manera: `<img href="archivo.jpg" style="anchura: 50%;" />`
* Cuando no se especifican atributos de anchura o altura, se recurre a la resolución de la imagen y a los metadatos de
  ppp incrustados en el archivo de imagen.


#### Notas a pie de página

Markdown de Pandoc permite notas a pie de página, utilizando la siguiente sintaxis:

```
He aquí una referencia a pie de página,[^1] y otra.[^longnote].

[^1]: Aquí está la nota a pie de página.

[^longnote]: Aquí hay una con varios bloques.

    Los párrafos subsiguientes están sangrados para mostrar que
pertenecen a la nota anterior.

        { some.code }

    Se puede sangrar todo el párrafo o sólo la primera
    línea. De este modo, las notas a pie de página de varios párrafos funcionan como
    elementos de lista multipárrafo.

Este párrafo no formará parte de la nota, porque
no está sangrado.
```

Los identificadores de las referencias de las notas a pie de página no pueden contener espacios, tabulaciones ni nuevas
líneas. Estos identificadores se utilizan únicamente para correlacionar la referencia de la nota a pie de página con la
propia nota; en la salida, las notas a pie de página se numerarán secuencialmente.

No es necesario que las notas a pie de página se sitúen al final del documento. Pueden aparecer en cualquier lugar,
excepto dentro de otros elementos del bloque (listas, citas en bloque, tablas, etc.). Cada nota a pie de página debe
estar separada del contenido circundante (incluidas otras notas) por líneas en blanco.

También se permiten las notas al pie en línea (aunque, a diferencia de las notas normales, no pueden contener varios
párrafos). La sintaxis es la siguiente:

```
Aquí tiene una nota inline.^[Las notas inline son más fáciles de escribir, ya que
es más fácil de escribir, ya que no hay que elegir un identificador y desplazarse hacia abajo para escribir la
nota].
```

Las notas a pie de página pueden mezclarse libremente con las notas a pie de página normales.


------------

## Referencias

* Este documento fue traducido y basado en el documento original del
  [sitio oficial de Pandoc](https://pandoc.org/MANUAL.html){target=_blank}.
* Si desea conocer mas acerca de esta fabulosa herramienta, consulte el
  [Manual de Pandoc](https://pandoc.org/MANUAL.html){target=_blank}.
* La traducción de este documento fue apoyada en parte por la herramienta web
  [DeepL](https://www.deepl.com/translator){target=_blank}
