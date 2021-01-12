import React from "react";
import { Link } from "gatsby";

import { rhythm, scale } from "../utils/typography";

const Layout = ({ location, title, children }) => {
    const rootPath = `${__PATH_PREFIX__}/`;
    let header;

    if (location.pathname === rootPath) {
        header = (
            <h1
                style={{
                    ...scale(1.5),
                    marginBottom: rhythm(1.5),
                    marginTop: 0
                }}
            >
                <Link
                    style={{
                        boxShadow: `none`,
                        color: `inherit`
                    }}
                    to={`/`}
                >
                    {title}
                </Link>
            </h1>
        );
    } else {
        header = (
            <h3
                style={{
                    fontFamily: `Montserrat, sans-serif`,
                    marginTop: 0
                }}
            >
                <Link
                    style={{
                        boxShadow: `none`,
                        color: `inherit`
                    }}
                    to={`/`}
                >
                    {title}
                </Link>
            </h3>
        );
    }
    return (
        <div
            style={{
                marginLeft: `auto`,
                marginRight: `auto`,
                maxWidth: rhythm(24),
                padding: `${rhythm(1.5)} ${rhythm(3 / 4)}`
            }}
        >
            <header>{header}</header>
            <main>{children}</main>
            <footer style={{ margin: "auto", textAlign: "center" }}>
                <div>
                    © {new Date().getFullYear()}, Built with
                    {` `}
                    <a href="https://www.gatsbyjs.org">Gatsby</a>
                </div>
                <div>
                    备案编号：
                    <a href="https://beian.miit.gov.cn">
                        粤 ICP 备 18141697 号
                    </a>
                </div>
                <div>信息产业部备案管理系统</div>
            </footer>
        </div>
    );
};

export default Layout;
