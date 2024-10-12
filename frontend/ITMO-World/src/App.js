import { useState, useEffect } from 'react';
import bridge from '@vkontakte/vk-bridge';
import { View, SplitLayout, SplitCol, ScreenSpinner } from '@vkontakte/vkui';
import { useActiveVkuiLocation } from '@vkontakte/vk-mini-apps-router';

import { HomePanel, TasksPanel, AchievementsPanel, LeaderboardPanel, StatisticsPanel } from './panels';
import { DEFAULT_VIEW_PANELS } from './routes';

export const App = () => {
  const { panel: activePanel = DEFAULT_VIEW_PANELS.HOME } = useActiveVkuiLocation();
  const [fetchedUser, setUser] = useState();
  const [popout, setPopout] = useState();

  useEffect(() => {
    async function fetchData() {
      const user = await bridge.send('VKWebAppGetUserInfo');
      setUser(user);
      setPopout(null);
    }
    fetchData();
  }, []);

  const goTo = (panel) => {
    // Ваш код для переключения панелей
  };

  return (
      <SplitLayout popout={popout}>
        <SplitCol>
          <View activePanel={activePanel}>
            <HomePanel id="home" fetchedUser={fetchedUser} goTo={goTo} />
            <TasksPanel id="tasks" goTo={goTo} />
            <AchievementsPanel id="achievements" goTo={goTo} />
            <LeaderboardPanel id="leaderboard" goTo={goTo} />
            <StatisticsPanel id="statistics" goTo={goTo} />
          </View>
        </SplitCol>
      </SplitLayout>
  );
};
