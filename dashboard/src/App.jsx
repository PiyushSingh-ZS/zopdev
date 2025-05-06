import { HashRouter, Route, Routes } from 'react-router-dom';
import { ErrorBoundary } from './components/organisms/ErrorBoundary';
import { Suspense } from 'react';
import { routes } from './routes';
import CompleteLoader from '@/components/atom/Loaders/CompleteLoader.jsx';
import TopBar from '@/components/molecule/TopBar/index.jsx';

function App() {
  return (
    <HashRouter>
      <ErrorBoundary>
        <Suspense fallback={<CompleteLoader />}>
          <TopBar />
          <Routes>
            {routes.map(({ path, component: Component }) => (
              <Route key={path} path={path} element={<Component />} />
            ))}
          </Routes>
        </Suspense>
      </ErrorBoundary>
    </HashRouter>
  );
}

export default App;
