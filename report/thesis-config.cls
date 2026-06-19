\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{thesis-config}[2026/06/19 GraduationThesisTemplate adapter]

\LoadClass[12pt,a4paper,twoside,openany]{book}

\makeatletter

\newcommand{\advisor}[1]{\gdef\@advisor{#1}}
\newcommand{\studentid}[1]{\gdef\@studentid{#1}}
\newcommand{\department}[1]{\gdef\@department{#1}}
\newcommand{\university}[1]{\gdef\@university{#1}}
\newcommand{\universitycity}[1]{\gdef\@universitycity{#1}}
\newcommand{\degree}[1]{\gdef\@degree{#1}}
\newcommand{\degreemonth}[1]{\gdef\@degreemonth{#1}}
\newcommand{\degreeyear}[1]{\gdef\@degreeyear{#1}}

\advisor{}
\studentid{}
\department{}
\university{ĐẠI HỌC BÁCH KHOA HÀ NỘI}
\universitycity{HÀ NỘI}
\degree{ĐỒ ÁN TỐT NGHIỆP}
\degreemonth{\the\month}
\degreeyear{\the\year}

\renewcommand{\maketitle}{%
  \begin{titlepage}
    \thispagestyle{empty}
    \begin{center}
      {\bfseries\fontsize{15pt}{18pt}\selectfont
        \MakeUppercase{\@university}\par}
      \vspace{0.4cm}
      {\bfseries\fontsize{15pt}{18pt}\selectfont
        \MakeUppercase{\@department}\par}

      \vspace{0.8cm}
      \includegraphics[width=2.43cm,height=3.94cm,keepaspectratio]
        {GraduationThesisTemplate-main/Images/Logo_Hust.png}

      \vspace{1.2cm}
      {\sffamily\bfseries\fontsize{21pt}{25pt}\selectfont
        \MakeUppercase{\@degree}\par}

      \vspace{1.2cm}
      {\sffamily\bfseries\fontsize{17pt}{22pt}\selectfont
        \MakeUppercase{\@title}\par}

      \vfill
      \begin{tabular}{@{}ll@{}}
        \textbf{Giảng viên hướng dẫn:} & \@advisor \\
        \textbf{Sinh viên thực hiện:} & \@author \\
        \textbf{Mã số sinh viên:} & \@studentid
      \end{tabular}

      \vfill
      {\bfseries
        \MakeUppercase{\@universitycity}, \@degreemonth/\@degreeyear}
    \end{center}
  \end{titlepage}
  \cleardoublepage
}

\makeatother
\endinput
