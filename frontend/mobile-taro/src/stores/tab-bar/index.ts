import { create } from 'zustand';
import type { TabBarState } from './type';

const useTabBarStore = create<TabBarState>((set) => ({
    selected: 0,
    setSelected: (index) => set({ selected: index }),
}));

export default useTabBarStore;
