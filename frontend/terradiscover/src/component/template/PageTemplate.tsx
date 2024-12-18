import { Outlet } from "react-router-dom"
import style from "./PageTemplate.module.css"

export const PageTemplate: React.FC = () => {
    return (
        <div className={style.page_template}>
            <Outlet />
        </div>
    )
}