// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(2)
    t := New(true, true)

    /*package main
      type Leg Peg {
       *Tree
      }*/
    t.AddPackage("main")
    t.AddLeg("Leg")
    t.AddState(`
 *Tree
`)
    t.AddYYSType("int")

    /* Grammar         <- - 'package' - Identifier      { p.AddPackage(buffer[begin:end]) }
       'type' - 'YYSTYPE' - Identifier { p.AddYYSType(buffer[begin:end]) } 
       'type' - Identifier         { p.AddLeg(buffer[begin:end]) }
       'Peg' - Action              { p.AddState(buffer[begin:end]) }
       ( Declaration | Definition)+ Trailer? EndOfFile */
    t.AddRule("Grammar")
    t.AddName("-")
    t.AddCharacter(`p`)
    t.AddCharacter(`a`)
    t.AddSequence()
    t.AddCharacter(`c`)
    t.AddSequence()
    t.AddCharacter(`k`)
    t.AddSequence()
    t.AddCharacter(`a`)
    t.AddSequence()
    t.AddCharacter(`g`)
    t.AddSequence()
    t.AddCharacter(`e`)
    t.AddSequence()
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddName("Identifier")
    t.AddSequence()
    t.AddAction(" p.AddPackage(buffer[begin:end]) ")
    t.AddSequence()
    t.AddCharacter(`Y`)
    t.AddCharacter(`Y`)
    t.AddSequence()
    t.AddCharacter(`S`)
    t.AddSequence()
    t.AddCharacter(`T`)
    t.AddSequence()
    t.AddCharacter(`Y`)
    t.AddSequence()
    t.AddCharacter(`P`)
    t.AddSequence()
    t.AddCharacter(`E`)
    t.AddSequence()
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddName("Identifier")
    t.AddSequence()
    t.AddAction(" p.AddYYSType(buffer[begin:end]) ")
    t.AddSequence()
    t.AddCharacter(`t`)
    t.AddCharacter(`y`)
    t.AddSequence()
    t.AddCharacter(`p`)
    t.AddSequence()
    t.AddCharacter(`e`)
    t.AddSequence()
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddName("Identifier")
    t.AddSequence()
    t.AddAction(" p.AddLeg(buffer[begin:end]) ")
    t.AddSequence()
    t.AddCharacter(`P`)
    t.AddCharacter(`e`)
    t.AddSequence()
    t.AddCharacter(`g`)
    t.AddSequence()
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddName("Action")
    t.AddSequence()
    t.AddAction(" p.AddState(buffer[begin:end]) ")
    t.AddSequence()
    t.AddName("Declaration")
    t.AddName("Definition")
    t.AddAlternate()
    t.AddPlus()
    t.AddSequence()
    t.AddName("Trailer")
    t.AddQuery()
    t.AddSequence()
    t.AddName("EndOfFile")
    t.AddSequence()
    t.AddExpression()

    /* Declaration = '%{' < ( !'%}' . )* >  {  p.AddDeclaration(buffer[begin:end])  }* RPERCENT */
    t.AddRule("Declaration")
    t.AddCharacter("%")
    t.AddCharacter("{")
    t.AddSequence()
    t.AddPush()
    t.AddCharacter("%")
    t.AddCharacter("}")
    t.AddSequence()
    t.AddPush()
    t.AddPeekNot()
    t.AddDot()
    t.AddSequence()
    t.AddStar()
    t.AddPush()
    t.AddSequence()
    t.AddName("RPERCENT")
    t.AddSequence()
    t.AddAction("  p.AddDeclaration(buffer[begin:end])  ")
    t.AddSequence()
    t.AddExpression()

    /* trailer =       '%%' < .* > { p.AddTrailer(buffer[begin:end]) }*/
    t.AddRule("Trailer")
    t.AddCharacter("%")
    t.AddCharacter("%")
    t.AddSequence()
    t.AddDot()
    t.AddStar()
    t.AddPush()
    t.AddAction(" p.AddTrailer(buffer[begin:end]) ")
    t.AddSequence()
    t.AddSequence()
    t.AddExpression()

    /* Definition      <- Identifier                   { p.AddRule(buffer[begin:end]) }
       Equal Expression         { p.AddExpression() }*/
    t.AddRule("Definition")
    t.AddName("Identifier")
    t.AddAction(" p.AddRule(buffer[begin:end]) ")
    t.AddSequence()
    t.AddName("Equal")
    t.AddSequence()
    t.AddName("Expression")
    t.AddSequence()
    t.AddAction(" p.AddExpression() ")
    t.AddSequence()
    // t.AddName("Identifier")
    // t.AddName("Equal")
    // t.AddSequence()
    // t.AddDot()
    // t.AddPeekNot()
    // t.AddAlternate()
    // t.AddPeekFor()
    // t.AddSequence()
    t.AddExpression()

    /* Expression      <- Sequence (Bar Sequence     { p.AddAlternate() }
               )* (Bar           { p.AddNil(); p.AddAlternate() }
                  )?
    / { p.AddNil() } */
    t.AddRule("Expression")
    t.AddName("Sequence")
    t.AddName("Bar")
    t.AddName("Sequence")
    t.AddSequence()
    t.AddAction(" p.AddAlternate() ")
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddName("Bar")
    t.AddAction(" p.AddNil(); p.AddAlternate() ")
    t.AddSequence()
    t.AddQuery()
    t.AddSequence()
    t.AddAction(" p.AddNil() ")
    t.AddAlternate()
    t.AddExpression()

    /* Sequence        <- Prefix (Prefix               { p.AddSequence() }
       )* */
    t.AddRule("Sequence")
    t.AddName("Prefix")
    t.AddName("Prefix")
    t.AddAction(" p.AddSequence() ")
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddExpression()

    /* Prefix          <- And Action                   { p.AddPredicate(buffer[begin:end]) }
       / And Suffix                   { p.AddPeekFor() }
       / Not Suffix                   { p.AddPeekNot() }
       /     Suffix */
    t.AddRule("Prefix")
    t.AddName("And")
    t.AddName("Action")
    t.AddSequence()
    t.AddAction(" p.AddPredicate(buffer[begin:end]) ")
    t.AddSequence()
    t.AddName("And")
    t.AddName("Suffix")
    t.AddSequence()
    t.AddAction(" p.AddPeekFor() ")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Not")
    t.AddName("Suffix")
    t.AddSequence()
    t.AddAction(" p.AddPeekNot() ")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Suffix")
    t.AddAlternate()
    t.AddExpression()

    /* Suffix          <- Primary (Question            { p.AddQuery() }
       / Star             { p.AddStar() }
       / Plus             { p.AddPlus() }
     )? */
    t.AddRule("Suffix")
    t.AddName("Primary")
    t.AddName("Question")
    t.AddAction(" p.AddQuery() ")
    t.AddSequence()
    t.AddName("Star")
    t.AddAction(" p.AddStar() ")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Plus")
    t.AddAction(" p.AddPlus() ")
    t.AddSequence()
    t.AddAlternate()
    t.AddQuery()
    t.AddSequence()
    t.AddExpression()

    /* Primary         = Identifier { p.AddVariable(buffer[begin:end]) } Colon Identifier !EQUAL { p.AddName(buffer[begin:end]) }
       / Identifier !Equal        { p.AddName(buffer[begin:end]) }
       / Open Expression Close
       / Literal
       / Class
       / Dot                          { p.AddDot() }
       / Action                       { p.AddAction(buffer[begin:end]) }
       / Begin Expression End         { p.AddPush() }*/
    t.AddRule("Primary")
    t.AddName("Identifier")
    t.AddAction(" p.AddVariable(buffer[begin:end]) ")
    t.AddSequence()
    t.AddName("Colon")
    t.AddSequence()
    t.AddName("Identifier")
    t.AddSequence()
    t.AddName("Equal")
    t.AddPeekNot()
    t.AddSequence()
    t.AddAction(" p.AddName(buffer[begin:end]) ")
    t.AddSequence()
    t.AddName("Identifier")
    t.AddName("Equal")
    t.AddPeekNot()
    t.AddSequence()
    t.AddAction(" p.AddName(buffer[begin:end]) ")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Open")
    t.AddName("Expression")
    t.AddSequence()
    t.AddName("Close")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Literal")
    t.AddAlternate()
    t.AddName("Class")
    t.AddAlternate()
    t.AddName("Dot")
    t.AddAction(" p.AddDot() ")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Action")
    t.AddAction(" p.AddAction(buffer[begin:end]) ")
    t.AddSequence()
    t.AddAlternate()
    t.AddName("Begin")
    t.AddName("Expression")
    t.AddSequence()
    t.AddName("End")
    t.AddSequence()
    t.AddAction(" p.AddPush() ")
    t.AddSequence()
    t.AddAlternate()
    t.AddExpression()

    /* Identifier  = < [-a-zA-Z_][-a-zA-Z_0-9]* > - */
    t.AddRule("Identifier")

    t.AddCharacter(`-`)
    t.AddCharacter(`a`)
    t.AddCharacter(`z`)
    t.AddDoubleRange()
    t.AddAlternate()
    t.AddCharacter(`_`)
    t.AddAlternate()
    t.AddCharacter(`-`)
    t.AddCharacter(`a`)
    t.AddCharacter(`z`)
    t.AddDoubleRange()
    t.AddAlternate()
    t.AddCharacter(`0`)
    t.AddCharacter(`9`)
    t.AddRange()
    t.AddAlternate()
    t.AddCharacter(`_`)
    t.AddAlternate()
    t.AddStar()
    t.AddSequence()
    t.AddPush()
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Literal         <- ['] (!['] Char)? (!['] Char          { p.AddSequence() }
                                           )* ['] -
                         / ["] (!["] DoubleChar)? (!["] DoubleChar          { p.AddSequence() }
                                                  )* ["] - */
    t.AddRule("Literal")
    t.AddCharacter(`'`)
    t.AddCharacter(`'`)
    t.AddPeekNot()
    t.AddName("Char")
    t.AddSequence()
    t.AddQuery()
    t.AddSequence()
    t.AddCharacter(`'`)
    t.AddPeekNot()
    t.AddName("Char")
    t.AddSequence()
    t.AddAction(` p.AddSequence() `)
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddCharacter(`'`)
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddCharacter(`"`)
    t.AddCharacter(`"`)
    t.AddPeekNot()
    t.AddName("DoubleChar")
    t.AddSequence()
    t.AddQuery()
    t.AddSequence()
    t.AddCharacter(`"`)
    t.AddPeekNot()
    t.AddName("DoubleChar")
    t.AddSequence()
    t.AddAction(` p.AddSequence() `)
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddCharacter(`"`)
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddAlternate()
    t.AddExpression()

    /* Class  <- ( '[[' ( '^' DoubleRanges              { p.AddPeekNot(); p.AddDot(); p.AddSequence() }
                            / DoubleRanges )?
                       ']]'
                     / '[' ( '^' Ranges                     { p.AddPeekNot(); p.AddDot(); p.AddSequence() }
                           / Ranges )?
                       ']' )
                     - */
    t.AddRule("Class")
    t.AddCharacter(`[`)
    t.AddCharacter(`[`)
    t.AddSequence()
    t.AddCharacter(`^`)
    t.AddName("DoubleRanges")
    t.AddSequence()
    t.AddAction(` p.AddPeekNot(); p.AddDot(); p.AddSequence() `)
    t.AddSequence()
    t.AddName("DoubleRanges")
    t.AddAlternate()
    t.AddQuery()
    t.AddSequence()
    t.AddCharacter(`]`)
    t.AddCharacter(`]`)
    t.AddSequence()
    t.AddSequence()
    t.AddCharacter(`[`)
    t.AddCharacter(`^`)
    t.AddName("Ranges")
    t.AddSequence()
    t.AddAction(` p.AddPeekNot(); p.AddDot(); p.AddSequence() `)
    t.AddSequence()
    t.AddName("Ranges")
    t.AddAlternate()
    t.AddQuery()
    t.AddSequence()
    t.AddCharacter(`]`)
    t.AddSequence()
    t.AddAlternate()
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Ranges          <- !']' Range (!']' Range  { p.AddAlternate() }
                                     )* */
    t.AddRule("Ranges")
    t.AddCharacter(`]`)
    t.AddPeekNot()
    t.AddName("Range")
    t.AddSequence()
    t.AddCharacter(`]`)
    t.AddPeekNot()
    t.AddName("Range")
    t.AddSequence()
    t.AddAction(" p.AddAlternate() ")
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddExpression()

    /* DoubleRanges          <- !']]' DoubleRange (!']]' DoubleRange  { p.AddAlternate() }
                                                  )* */
    t.AddRule("DoubleRanges")
    t.AddCharacter(`]`)
    t.AddCharacter(`]`)
    t.AddSequence()
    t.AddPeekNot()
    t.AddName("DoubleRange")
    t.AddSequence()
    t.AddCharacter(`]`)
    t.AddCharacter(`]`)
    t.AddSequence()
    t.AddPeekNot()
    t.AddName("DoubleRange")
    t.AddSequence()
    t.AddAction(" p.AddAlternate() ")
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddExpression()

    /* Range           <- Char '-' Char { p.AddRange() }
                            / Char */
    t.AddRule("Range")
    t.AddName("Char")
    t.AddCharacter(`-`)
    t.AddSequence()
    t.AddName("Char")
    t.AddSequence()
    t.AddAction(" p.AddRange() ")
    t.AddSequence()
    t.AddName("Char")
    t.AddAlternate()
    t.AddExpression()

    /* DoubleRange      <- Char '-' Char { p.AddDoubleRange() }
                             / DoubleChar */
    t.AddRule("DoubleRange")
    t.AddName("Char")
    t.AddCharacter(`-`)
    t.AddSequence()
    t.AddName("Char")
    t.AddSequence()
    t.AddAction(" p.AddDoubleRange() ")
    t.AddSequence()
    t.AddName("DoubleChar")
    t.AddAlternate()
    t.AddExpression()

    /* Char            <- Escape
                            / !'\\' <.>                  { p.AddCharacter(buffer[begin:end]) } */
    t.AddRule("Char")
    t.AddName("Escape")
    t.AddCharacter("\\")
    t.AddPeekNot()
    t.AddDot()
    t.AddPush()
    t.AddSequence()
    t.AddAction(` p.AddCharacter(buffer[begin:end]) `)
    t.AddSequence()
    t.AddAlternate()
    t.AddExpression()

    /* DoubleChar      <- Escape
                            / <[a-zA-Z]>                 { p.AddDoubleCharacter(buffer[begin:end]) }
                            / !'\\' <.>                  { p.AddCharacter(buffer[begin:end]) } */
    t.AddRule("DoubleChar")
    t.AddName("Escape")
    t.AddCharacter(`a`)
    t.AddCharacter(`z`)
    t.AddRange()
    t.AddCharacter(`A`)
    t.AddCharacter(`Z`)
    t.AddRange()
    t.AddAlternate()
    t.AddPush()
    t.AddAction(` p.AddDoubleCharacter(buffer[begin:end]) `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddPeekNot()
    t.AddDot()
    t.AddPush()
    t.AddSequence()
    t.AddAction(` p.AddCharacter(buffer[begin:end]) `)
    t.AddSequence()
    t.AddAlternate()
    t.AddExpression()

    /* Escape            <- "\\a"                      { p.AddCharacter("\a") }   # bell
                          / "\\b"                      { p.AddCharacter("\b") }   # bs
                              / "\\e"                      { p.AddCharacter("\x1B") } # esc
                              / "\\f"                      { p.AddCharacter("\f") }   # ff
                              / "\\n"                      { p.AddCharacter("\n") }   # nl
                              / "\\r"                      { p.AddCharacter("\r") }   # cr
                              / "\\t"                      { p.AddCharacter("\t") }   # ht
                              / "\\v"                      { p.AddCharacter("\v") }   # vt
                              / "\\'"                      { p.AddCharacter("'") }
                              / '\\"'                      { p.AddCharacter("\"") }
                              / '\\['                      { p.AddCharacter("[") }
                              / '\\]'                      { p.AddCharacter("]") }
                              / '\\-'                      { p.AddCharacter("-") }
                              / '\\' <[0-3][0-7][0-7]>     { p.AddOctalCharacter(buffer[begin:end]) }
                              / '\\' <[0-7][0-7]?>         { p.AddOctalCharacter(buffer[begin:end]) }
                              / '\\\\'                     { p.AddCharacter("\\") } */
    t.AddRule("Escape")
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`a`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\a") `)
    t.AddSequence()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`b`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\b") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`e`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\x1B") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`f`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\f") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`n`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\n") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`r`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\r") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`t`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\t") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddDoubleCharacter(`v`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\v") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`'`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("'") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`"`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\"") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`[`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("[") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`]`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("]") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`-`)
    t.AddSequence()
    t.AddAction(` p.AddCharacter("-") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`0`)
    t.AddCharacter(`3`)
    t.AddRange()
    t.AddCharacter(`0`)
    t.AddCharacter(`7`)
    t.AddRange()
    t.AddSequence()
    t.AddCharacter(`0`)
    t.AddCharacter(`7`)
    t.AddRange()
    t.AddSequence()
    t.AddPush()
    t.AddSequence()
    t.AddAction(` p.AddOctalCharacter(buffer[begin:end]) `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter(`0`)
    t.AddCharacter(`7`)
    t.AddRange()
    t.AddCharacter(`0`)
    t.AddCharacter(`7`)
    t.AddRange()
    t.AddQuery()
    t.AddSequence()
    t.AddPush()
    t.AddSequence()
    t.AddAction(` p.AddOctalCharacter(buffer[begin:end]) `)
    t.AddSequence()
    t.AddAlternate()
    t.AddCharacter("\\")
    t.AddCharacter("\\")
    t.AddSequence()
    t.AddAction(` p.AddCharacter("\\") `)
    t.AddSequence()
    t.AddAlternate()
    t.AddExpression()

    /* Action          <- '{' < Braces* > '}' - */
    t.AddRule("Action")
    t.AddCharacter(`{`)
    t.AddName("Braces")
    t.AddStar()
    t.AddPush()
    t.AddSequence()
    t.AddCharacter(`}`)
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Braces <-        '{' Braces* '}' /   !'}' . */
    t.AddRule("Braces")
    t.AddCharacter(`{`)
    t.AddName("Braces")
    t.AddStar()
    t.AddSequence()
    t.AddCharacter(`}`)
    t.AddSequence()
    t.AddCharacter(`}`)
    t.AddPeekNot()
    t.AddDot()
    t.AddSequence()
    t.AddAlternate()
    t.AddExpression()

    /* Equal       <- '=' - */
    t.AddRule("Equal")
    t.AddCharacter(`=`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Colon      <- ':' - */
    t.AddRule("Colon")
    t.AddCharacter(`:`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Bar           <- '|' - */
    t.AddRule("Bar")
    t.AddCharacter(`|`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* And             <- '&' - */
    t.AddRule("And")
    t.AddCharacter(`&`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Not             <- '!' - */
    t.AddRule("Not")
    t.AddCharacter(`!`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Question        <- '?' - */
    t.AddRule("Question")
    t.AddCharacter(`?`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Star            <- '*' - */
    t.AddRule("Star")
    t.AddCharacter(`*`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Plus            <- '+' - */
    t.AddRule("Plus")
    t.AddCharacter(`+`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Open            <- '(' - */
    t.AddRule("Open")
    t.AddCharacter(`(`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Close           <- ')' - */
    t.AddRule("Close")
    t.AddCharacter(`)`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* Dot             <- '.' - */
    t.AddRule("Dot")
    t.AddCharacter(`.`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* RPERCENT =      '%}' - */
    t.AddRule("RPERCENT")
    t.AddCharacter("%")
    t.AddCharacter("}")
    t.AddSequence()
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* -         <- (Space / Comment)* */
    t.AddRule("-")
    t.AddName("Space")
    t.AddName("Comment")
    t.AddAlternate()
    t.AddStar()
    t.AddExpression()

    /* Comment         <- '#' (!EndOfLine .)* EndOfLine */
    t.AddRule("Comment")
    t.AddCharacter(`#`)
    t.AddName("EndOfLine")
    t.AddPeekNot()
    t.AddDot()
    t.AddSequence()
    t.AddStar()
    t.AddSequence()
    t.AddName("EndOfLine")
    t.AddSequence()
    t.AddExpression()

    /* Space           <- ' ' / '\t' / EndOfLine */
    t.AddRule("Space")
    t.AddCharacter(` `)
    t.AddCharacter("\t")
    t.AddAlternate()
    t.AddName("EndOfLine")
    t.AddAlternate()
    t.AddExpression()

    /* EndOfLine       <- '\r\n' / '\n' / '\r' */
    t.AddRule("EndOfLine")
    t.AddCharacter("\r")
    t.AddCharacter("\n")
    t.AddSequence()
    t.AddCharacter("\n")
    t.AddAlternate()
    t.AddCharacter("\r")
    t.AddAlternate()
    t.AddExpression()

    /* EndOfFile       <- !. */
    t.AddRule("EndOfFile")
    t.AddDot()
    t.AddPeekNot()
    t.AddExpression()

    /* Begin           <- '<' - */
    t.AddRule("Begin")
    t.AddCharacter(`<`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    /* End             <- '>' - */
    t.AddRule("End")
    t.AddCharacter(`>`)
    t.AddName("-")
    t.AddSequence()
    t.AddExpression()

    t.Compile("bootstrap.leg.go")
}