/**
 * TabBar Store 相关的 TypeScript 类型定义
 */
export interface TabBarState {
    /**
     * 当前选中的 Tab 索引
     */
    selected: number;
    /**
     * 设置选中的 Tab 索引
     * @param index 索引值
     */
    setSelected: (index: number) => void;
}
