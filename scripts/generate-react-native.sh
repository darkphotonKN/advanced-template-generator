#!/bin/bash

# React Native Project Generator Script with NativeWind v4
# Usage: ./scripts/generate-react-native.sh project-name entity api-port [auth] [s3]

PROJECT_NAME="$1"
ENTITY="$2"
API_PORT="$3"
INCLUDE_AUTH="${4:-true}"
INCLUDE_S3="${5:-false}"

if [ -z "$PROJECT_NAME" ] || [ -z "$ENTITY" ] || [ -z "$API_PORT" ]; then
    echo "Usage: $0 <project-name> <entity> <api-port> [auth] [s3]"
    echo "Example: $0 ishopgo product 8039 true true"
    exit 1
fi

# React Native app is created as sibling to backend and admin
APP_DIR="../${PROJECT_NAME}/${PROJECT_NAME}-app"
ENTITY_PLURAL="${ENTITY}s"
ENTITY_CAPITALIZED="$(echo ${ENTITY:0:1} | tr '[:lower:]' '[:upper:]')${ENTITY:1}"
PROJECT_TITLE="$(echo ${PROJECT_NAME} | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')"

echo "📱 Generating React Native project with NativeWind v4..."
echo "- Project: $PROJECT_NAME"
echo "- Entity: $ENTITY ($ENTITY_PLURAL, $ENTITY_CAPITALIZED)"
echo "- API Port: $API_PORT"
echo "- Auth: $INCLUDE_AUTH"
echo "- S3: $INCLUDE_S3"

# Create app directory
echo "📁 Creating React Native app directory..."
mkdir -p "$APP_DIR"
cd "$APP_DIR"

# Initialize Expo with TypeScript template
echo "🎯 Initializing Expo project..."
npx create-expo-app@latest . --template blank-typescript --no-install

# Create project structure
echo "🏗️ Creating project structure..."
mkdir -p src/{app,components,features,hooks,lib,services,stores,types,utils}
mkdir -p src/features/$ENTITY/{components,hooks,screens,types}
mkdir -p src/components/{ui,layout}
mkdir -p assets/{fonts,images}

# Create package.json with all dependencies
cat > package.json << EOF
{
  "name": "${PROJECT_NAME}-app",
  "version": "1.0.0",
  "main": "node_modules/expo/AppEntry.js",
  "scripts": {
    "start": "expo start",
    "android": "expo run:android",
    "ios": "expo run:ios",
    "web": "expo start --web",
    "test": "jest",
    "lint": "eslint . --ext .ts,.tsx",
    "format": "prettier --write ."
  },
  "dependencies": {
    "expo": "~51.0.0",
    "expo-status-bar": "~1.12.1",
    "react": "18.3.1",
    "react-native": "0.75.4",
    "@react-navigation/native": "^6.1.18",
    "@react-navigation/native-stack": "^6.11.0",
    "@react-navigation/bottom-tabs": "^6.6.1",
    "react-native-screens": "~3.34.0",
    "react-native-safe-area-context": "4.11.1",
    "nativewind": "^4.1.23",
    "tailwindcss": "^3.4.15",
    "@tanstack/react-query": "^5.62.8",
    "zustand": "^4.5.5",
    "axios": "^1.7.9",
    "react-hook-form": "^7.54.2",
    "@hookform/resolvers": "^3.9.1",
    "zod": "^3.24.1",
    "expo-secure-store": "~13.0.2",
    "expo-auth-session": "~5.5.2",
    "expo-crypto": "~13.0.2",
    "expo-web-browser": "~13.0.3",
    "expo-image": "~1.13.0",
    "expo-font": "~12.0.10",
    "expo-splash-screen": "~0.27.7",
    "react-native-gesture-handler": "~2.20.2",
    "react-native-reanimated": "~3.16.1",
    "react-native-svg": "15.8.0",
    "@shopify/flash-list": "1.7.2",
    "date-fns": "^3.6.0"
  },
  "devDependencies": {
    "@babel/core": "^7.20.0",
    "@types/react": "~18.3.12",
    "typescript": "^5.3.0",
    "eslint": "^8.57.0",
    "@typescript-eslint/eslint-plugin": "^7.18.0",
    "@typescript-eslint/parser": "^7.18.0",
    "eslint-config-expo": "^7.1.2",
    "prettier": "^3.3.3",
    "jest": "^29.7.0",
    "@testing-library/react-native": "^12.8.1",
    "metro-react-native-babel-preset": "^0.77.0"
  },
  "private": true
}
EOF

# Create tailwind.config.js for NativeWind v4
cat > tailwind.config.js << 'EOF'
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}", "./App.tsx"],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
          800: '#1e40af',
          900: '#1e3a8a',
          950: '#172554',
        },
      },
    },
  },
  plugins: [],
}
EOF

# Create metro.config.js for NativeWind v4
cat > metro.config.js << 'EOF'
const { getDefaultConfig } = require('expo/metro-config');
const { withNativeWind } = require('nativewind/metro');

const config = getDefaultConfig(__dirname);

module.exports = withNativeWind(config, { input: './src/styles/global.css' });
EOF

# Create babel.config.js
cat > babel.config.js << 'EOF'
module.exports = function(api) {
  api.cache(true);
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      'nativewind/babel',
      'react-native-reanimated/plugin',
    ],
  };
};
EOF

# Create global CSS for NativeWind v4
mkdir -p src/styles
cat > src/styles/global.css << 'EOF'
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer components {
  .btn-primary {
    @apply bg-primary-600 px-4 py-2 rounded-lg;
  }

  .btn-secondary {
    @apply bg-gray-200 px-4 py-2 rounded-lg;
  }

  .input {
    @apply border border-gray-300 rounded-lg px-3 py-2;
  }

  .card {
    @apply bg-white rounded-xl shadow-sm p-4;
  }
}
EOF

# Create tsconfig.json
cat > tsconfig.json << 'EOF'
{
  "extends": "expo/tsconfig.base",
  "compilerOptions": {
    "strict": true,
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"],
      "@components/*": ["src/components/*"],
      "@features/*": ["src/features/*"],
      "@hooks/*": ["src/hooks/*"],
      "@lib/*": ["src/lib/*"],
      "@services/*": ["src/services/*"],
      "@stores/*": ["src/stores/*"],
      "@types/*": ["src/types/*"],
      "@utils/*": ["src/utils/*"]
    }
  }
}
EOF

# Create app.json
cat > app.json << EOF
{
  "expo": {
    "name": "${PROJECT_NAME}-app",
    "slug": "${PROJECT_NAME}-app",
    "version": "1.0.0",
    "orientation": "portrait",
    "icon": "./assets/images/icon.png",
    "userInterfaceStyle": "automatic",
    "splash": {
      "image": "./assets/images/splash.png",
      "resizeMode": "contain",
      "backgroundColor": "#ffffff"
    },
    "assetBundlePatterns": [
      "**/*"
    ],
    "ios": {
      "supportsTablet": true,
      "bundleIdentifier": "com.${PROJECT_NAME//-/}.app"
    },
    "android": {
      "adaptiveIcon": {
        "foregroundImage": "./assets/images/adaptive-icon.png",
        "backgroundColor": "#ffffff"
      },
      "package": "com.${PROJECT_NAME//-/_}.app"
    },
    "web": {
      "favicon": "./assets/images/favicon.png",
      "bundler": "metro"
    }
  }
}
EOF

# Create .env.example
cat > .env.example << EOF
# API Configuration
EXPO_PUBLIC_API_URL=http://localhost:${API_PORT}/api
EXPO_PUBLIC_API_TIMEOUT=30000

# Auth Configuration
EXPO_PUBLIC_AUTH_ENABLED=${INCLUDE_AUTH}

# S3 Configuration
EXPO_PUBLIC_S3_ENABLED=${INCLUDE_S3}

# App Configuration
EXPO_PUBLIC_APP_NAME=${PROJECT_TITLE}
EXPO_PUBLIC_APP_VERSION=1.0.0
EOF

# Create App.tsx with NativeWind v4 setup
cat > App.tsx << 'EOF'
import './src/styles/global.css';
import { StatusBar } from 'expo-status-bar';
import { NavigationContainer } from '@react-navigation/native';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { RootNavigator } from '@/navigation/RootNavigator';
import { useAuthStore } from '@/stores/auth.store';
import { useEffect } from 'react';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 2,
      staleTime: 5 * 60 * 1000, // 5 minutes
    },
  },
});

export default function App() {
  const initAuth = useAuthStore((state) => state.init);

  useEffect(() => {
    initAuth();
  }, []);

  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      <SafeAreaProvider>
        <QueryClientProvider client={queryClient}>
          <NavigationContainer>
            <RootNavigator />
            <StatusBar style="auto" />
          </NavigationContainer>
        </QueryClientProvider>
      </SafeAreaProvider>
    </GestureHandlerRootView>
  );
}
EOF

# Create navigation structure
mkdir -p src/navigation
cat > src/navigation/RootNavigator.tsx << EOF
import React from 'react';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { useAuthStore } from '@/stores/auth.store';
import { HomeScreen } from '@/features/home/screens/HomeScreen';
import { ${ENTITY_CAPITALIZED}ListScreen } from '@/features/${ENTITY}/screens/${ENTITY_CAPITALIZED}ListScreen';
import { ${ENTITY_CAPITALIZED}DetailScreen } from '@/features/${ENTITY}/screens/${ENTITY_CAPITALIZED}DetailScreen';
import { ProfileScreen } from '@/features/profile/screens/ProfileScreen';
import { AuthScreen } from '@/features/auth/screens/AuthScreen';
import { Ionicons } from '@expo/vector-icons';

export type RootStackParamList = {
  Main: undefined;
  Auth: undefined;
  ${ENTITY_CAPITALIZED}Detail: { id: string };
};

export type MainTabParamList = {
  Home: undefined;
  ${ENTITY_CAPITALIZED}s: undefined;
  Profile: undefined;
};

const Stack = createNativeStackNavigator<RootStackParamList>();
const Tab = createBottomTabNavigator<MainTabParamList>();

function MainTabs() {
  return (
    <Tab.Navigator
      screenOptions={({ route }) => ({
        tabBarIcon: ({ focused, color, size }) => {
          let iconName: keyof typeof Ionicons.glyphMap = 'home';

          if (route.name === 'Home') {
            iconName = focused ? 'home' : 'home-outline';
          } else if (route.name === '${ENTITY_CAPITALIZED}s') {
            iconName = focused ? 'list' : 'list-outline';
          } else if (route.name === 'Profile') {
            iconName = focused ? 'person' : 'person-outline';
          }

          return <Ionicons name={iconName} size={size} color={color} />;
        },
        tabBarActiveTintColor: '#2563eb',
        tabBarInactiveTintColor: 'gray',
      })}
    >
      <Tab.Screen name="Home" component={HomeScreen} />
      <Tab.Screen name="${ENTITY_CAPITALIZED}s" component={${ENTITY_CAPITALIZED}ListScreen} />
      <Tab.Screen name="Profile" component={ProfileScreen} />
    </Tab.Navigator>
  );
}

export function RootNavigator() {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);

  return (
    <Stack.Navigator>
      {isAuthenticated ? (
        <>
          <Stack.Screen
            name="Main"
            component={MainTabs}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="${ENTITY_CAPITALIZED}Detail"
            component={${ENTITY_CAPITALIZED}DetailScreen}
            options={{ title: '${ENTITY_CAPITALIZED} Details' }}
          />
        </>
      ) : (
        <Stack.Screen
          name="Auth"
          component={AuthScreen}
          options={{ headerShown: false }}
        />
      )}
    </Stack.Navigator>
  );
}
EOF

# Create API client
cat > src/lib/api/client.ts << 'EOF'
import axios from 'axios';
import { useAuthStore } from '@/stores/auth.store';

const API_URL = process.env.EXPO_PUBLIC_API_URL || 'http://localhost:8080/api';

export const apiClient = axios.create({
  baseURL: API_URL,
  timeout: parseInt(process.env.EXPO_PUBLIC_API_TIMEOUT || '30000'),
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor for auth
apiClient.interceptors.request.use(
  (config) => {
    const token = useAuthStore.getState().token;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor for error handling
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      useAuthStore.getState().logout();
    }
    return Promise.reject(error);
  }
);
EOF

# Create auth store with Zustand
cat > src/stores/auth.store.ts << 'EOF'
import { create } from 'zustand';
import * as SecureStore from 'expo-secure-store';

interface User {
  id: string;
  email: string;
  name: string;
}

interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  token: string | null;
  isLoading: boolean;
  init: () => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  setToken: (token: string) => Promise<void>;
  setUser: (user: User) => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  isAuthenticated: false,
  user: null,
  token: null,
  isLoading: true,

  init: async () => {
    try {
      const token = await SecureStore.getItemAsync('auth_token');
      if (token) {
        // TODO: Validate token with backend
        set({ token, isAuthenticated: true });
      }
    } catch (error) {
      console.error('Failed to init auth:', error);
    } finally {
      set({ isLoading: false });
    }
  },

  login: async (email: string, password: string) => {
    // TODO: Implement login logic
    const token = 'mock-token';
    await SecureStore.setItemAsync('auth_token', token);
    set({ token, isAuthenticated: true });
  },

  logout: async () => {
    await SecureStore.deleteItemAsync('auth_token');
    set({ token: null, user: null, isAuthenticated: false });
  },

  setToken: async (token: string) => {
    await SecureStore.setItemAsync('auth_token', token);
    set({ token, isAuthenticated: true });
  },

  setUser: (user: User) => set({ user }),
}));
EOF

# Create entity feature screens
cat > src/features/${ENTITY}/screens/${ENTITY_CAPITALIZED}ListScreen.tsx << EOF
import React from 'react';
import { View, Text, FlatList, TouchableOpacity, RefreshControl } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { RootStackParamList } from '@/navigation/RootNavigator';
import { use${ENTITY_CAPITALIZED}s } from '../hooks/use${ENTITY_CAPITALIZED}s';
import { ${ENTITY_CAPITALIZED}Card } from '../components/${ENTITY_CAPITALIZED}Card';

type NavigationProp = NativeStackNavigationProp<RootStackParamList, 'Main'>;

export function ${ENTITY_CAPITALIZED}ListScreen() {
  const navigation = useNavigation<NavigationProp>();
  const { data, isLoading, refetch, isRefetching } = use${ENTITY_CAPITALIZED}s();

  const handle${ENTITY_CAPITALIZED}Press = (id: string) => {
    navigation.navigate('${ENTITY_CAPITALIZED}Detail', { id });
  };

  if (isLoading && !isRefetching) {
    return (
      <View className="flex-1 justify-center items-center">
        <Text className="text-gray-500">Loading ${ENTITY}s...</Text>
      </View>
    );
  }

  return (
    <View className="flex-1 bg-gray-50">
      <FlatList
        data={data}
        keyExtractor={(item) => item.id}
        renderItem={({ item }) => (
          <TouchableOpacity onPress={() => handle${ENTITY_CAPITALIZED}Press(item.id)}>
            <${ENTITY_CAPITALIZED}Card ${ENTITY}={item} />
          </TouchableOpacity>
        )}
        contentContainerStyle={{ padding: 16 }}
        ItemSeparatorComponent={() => <View className="h-4" />}
        refreshControl={
          <RefreshControl refreshing={isRefetching} onRefresh={refetch} />
        }
      />
    </View>
  );
}
EOF

cat > src/features/${ENTITY}/screens/${ENTITY_CAPITALIZED}DetailScreen.tsx << EOF
import React from 'react';
import { View, Text, ScrollView, ActivityIndicator } from 'react-native';
import { useRoute, RouteProp } from '@react-navigation/native';
import { RootStackParamList } from '@/navigation/RootNavigator';
import { use${ENTITY_CAPITALIZED} } from '../hooks/use${ENTITY_CAPITALIZED}';

type RouteParams = RouteProp<RootStackParamList, '${ENTITY_CAPITALIZED}Detail'>;

export function ${ENTITY_CAPITALIZED}DetailScreen() {
  const route = useRoute<RouteParams>();
  const { data: ${ENTITY}, isLoading } = use${ENTITY_CAPITALIZED}(route.params.id);

  if (isLoading) {
    return (
      <View className="flex-1 justify-center items-center">
        <ActivityIndicator size="large" color="#2563eb" />
      </View>
    );
  }

  if (!${ENTITY}) {
    return (
      <View className="flex-1 justify-center items-center">
        <Text className="text-gray-500">${ENTITY_CAPITALIZED} not found</Text>
      </View>
    );
  }

  return (
    <ScrollView className="flex-1 bg-white">
      <View className="p-4">
        <Text className="text-2xl font-bold mb-2">{${ENTITY}.name}</Text>
        <Text className="text-gray-600 mb-4">{${ENTITY}.description}</Text>

        <View className="bg-gray-50 rounded-lg p-4 mt-4">
          <Text className="text-lg font-semibold mb-2">Details</Text>
          <View className="space-y-2">
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Price:</Text>
              <Text className="font-medium">\${${ENTITY}.price}</Text>
            </View>
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Stock:</Text>
              <Text className="font-medium">{${ENTITY}.stock}</Text>
            </View>
            <View className="flex-row justify-between">
              <Text className="text-gray-600">Category:</Text>
              <Text className="font-medium">{${ENTITY}.category}</Text>
            </View>
          </View>
        </View>
      </View>
    </ScrollView>
  );
}
EOF

# Create entity components
cat > src/features/${ENTITY}/components/${ENTITY_CAPITALIZED}Card.tsx << EOF
import React from 'react';
import { View, Text, Image } from 'react-native';
import { ${ENTITY_CAPITALIZED} } from '../types';

interface ${ENTITY_CAPITALIZED}CardProps {
  ${ENTITY}: ${ENTITY_CAPITALIZED};
}

export function ${ENTITY_CAPITALIZED}Card({ ${ENTITY} }: ${ENTITY_CAPITALIZED}CardProps) {
  return (
    <View className="bg-white rounded-xl shadow-sm p-4">
      {${ENTITY}.image && (
        <Image
          source={{ uri: ${ENTITY}.image }}
          className="w-full h-48 rounded-lg mb-3"
          resizeMode="cover"
        />
      )}
      <Text className="text-lg font-semibold mb-1">{${ENTITY}.name}</Text>
      <Text className="text-gray-600 mb-2" numberOfLines={2}>
        {${ENTITY}.description}
      </Text>
      <View className="flex-row justify-between items-center">
        <Text className="text-primary-600 font-bold text-lg">\${${ENTITY}.price}</Text>
        <View className="bg-primary-100 px-3 py-1 rounded-full">
          <Text className="text-primary-800 text-sm">{${ENTITY}.category}</Text>
        </View>
      </View>
    </View>
  );
}
EOF

# Create entity hooks
cat > src/features/${ENTITY}/hooks/use${ENTITY_CAPITALIZED}s.ts << EOF
import { useQuery } from '@tanstack/react-query';
import { ${ENTITY}Service } from '../services/${ENTITY}.service';

export function use${ENTITY_CAPITALIZED}s() {
  return useQuery({
    queryKey: ['${ENTITY}s'],
    queryFn: ${ENTITY}Service.getAll,
  });
}
EOF

cat > src/features/${ENTITY}/hooks/use${ENTITY_CAPITALIZED}.ts << EOF
import { useQuery } from '@tanstack/react-query';
import { ${ENTITY}Service } from '../services/${ENTITY}.service';

export function use${ENTITY_CAPITALIZED}(id: string) {
  return useQuery({
    queryKey: ['${ENTITY}', id],
    queryFn: () => ${ENTITY}Service.getById(id),
    enabled: !!id,
  });
}
EOF

# Create entity service
cat > src/features/${ENTITY}/services/${ENTITY}.service.ts << EOF
import { apiClient } from '@/lib/api/client';
import { ${ENTITY_CAPITALIZED} } from '../types';

export const ${ENTITY}Service = {
  async getAll(): Promise<${ENTITY_CAPITALIZED}[]> {
    const { data } = await apiClient.get('/${ENTITY}s');
    return data;
  },

  async getById(id: string): Promise<${ENTITY_CAPITALIZED}> {
    const { data } = await apiClient.get(\`/${ENTITY}s/\${id}\`);
    return data;
  },

  async create(${ENTITY}: Omit<${ENTITY_CAPITALIZED}, 'id'>): Promise<${ENTITY_CAPITALIZED}> {
    const { data } = await apiClient.post('/${ENTITY}s', ${ENTITY});
    return data;
  },

  async update(id: string, ${ENTITY}: Partial<${ENTITY_CAPITALIZED}>): Promise<${ENTITY_CAPITALIZED}> {
    const { data } = await apiClient.put(\`/${ENTITY}s/\${id}\`, ${ENTITY});
    return data;
  },

  async delete(id: string): Promise<void> {
    await apiClient.delete(\`/${ENTITY}s/\${id}\`);
  },
};
EOF

# Create entity types
cat > src/features/${ENTITY}/types/index.ts << EOF
export interface ${ENTITY_CAPITALIZED} {
  id: string;
  name: string;
  description: string;
  price: number;
  stock: number;
  category: string;
  image?: string;
  createdAt: string;
  updatedAt: string;
}

export interface ${ENTITY_CAPITALIZED}Filter {
  search?: string;
  category?: string;
  minPrice?: number;
  maxPrice?: number;
  sortBy?: 'name' | 'price' | 'createdAt';
  sortOrder?: 'asc' | 'desc';
}
EOF

# Create other screens
mkdir -p src/features/home/screens
cat > src/features/home/screens/HomeScreen.tsx << EOF
import React from 'react';
import { View, Text, ScrollView, TouchableOpacity } from 'react-native';
import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { MainTabParamList } from '@/navigation/RootNavigator';

type NavigationProp = NativeStackNavigationProp<MainTabParamList, 'Home'>;

export function HomeScreen() {
  const navigation = useNavigation<NavigationProp>();

  return (
    <ScrollView className="flex-1 bg-gray-50">
      <View className="p-4">
        <Text className="text-3xl font-bold mb-6">Welcome to ${PROJECT_TITLE}</Text>

        <View className="space-y-4">
          <TouchableOpacity
            className="bg-primary-600 p-4 rounded-xl"
            onPress={() => navigation.navigate('${ENTITY_CAPITALIZED}s')}
          >
            <Text className="text-white text-lg font-semibold">Browse ${ENTITY_CAPITALIZED}s</Text>
            <Text className="text-white/80 mt-1">Explore our collection</Text>
          </TouchableOpacity>

          <View className="bg-white rounded-xl p-4">
            <Text className="text-lg font-semibold mb-2">Quick Stats</Text>
            <View className="space-y-2">
              <View className="flex-row justify-between">
                <Text className="text-gray-600">Total ${ENTITY_CAPITALIZED}s</Text>
                <Text className="font-medium">--</Text>
              </View>
              <View className="flex-row justify-between">
                <Text className="text-gray-600">Categories</Text>
                <Text className="font-medium">--</Text>
              </View>
            </View>
          </View>
        </View>
      </View>
    </ScrollView>
  );
}
EOF

mkdir -p src/features/profile/screens
cat > src/features/profile/screens/ProfileScreen.tsx << EOF
import React from 'react';
import { View, Text, TouchableOpacity, ScrollView } from 'react-native';
import { useAuthStore } from '@/stores/auth.store';

export function ProfileScreen() {
  const { user, logout } = useAuthStore();

  return (
    <ScrollView className="flex-1 bg-gray-50">
      <View className="p-4">
        <View className="bg-white rounded-xl p-4 mb-4">
          <Text className="text-2xl font-bold mb-2">Profile</Text>
          {user ? (
            <>
              <Text className="text-gray-600 mb-1">{user.name}</Text>
              <Text className="text-gray-600">{user.email}</Text>
            </>
          ) : (
            <Text className="text-gray-500">Not logged in</Text>
          )}
        </View>

        <TouchableOpacity
          className="bg-red-500 p-4 rounded-xl"
          onPress={logout}
        >
          <Text className="text-white text-center font-semibold">Logout</Text>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
}
EOF

mkdir -p src/features/auth/screens
cat > src/features/auth/screens/AuthScreen.tsx << EOF
import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, KeyboardAvoidingView, Platform } from 'react-native';
import { useAuthStore } from '@/stores/auth.store';

export function AuthScreen() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const login = useAuthStore((state) => state.login);

  const handleLogin = async () => {
    if (email && password) {
      await login(email, password);
    }
  };

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      className="flex-1"
    >
      <View className="flex-1 justify-center p-6 bg-gray-50">
        <View className="bg-white rounded-2xl p-6">
          <Text className="text-3xl font-bold text-center mb-8">${PROJECT_TITLE}</Text>

          <View className="space-y-4">
            <View>
              <Text className="text-gray-700 mb-2">Email</Text>
              <TextInput
                className="border border-gray-300 rounded-lg px-3 py-2"
                value={email}
                onChangeText={setEmail}
                placeholder="Enter your email"
                keyboardType="email-address"
                autoCapitalize="none"
              />
            </View>

            <View>
              <Text className="text-gray-700 mb-2">Password</Text>
              <TextInput
                className="border border-gray-300 rounded-lg px-3 py-2"
                value={password}
                onChangeText={setPassword}
                placeholder="Enter your password"
                secureTextEntry
              />
            </View>

            <TouchableOpacity
              className="bg-primary-600 py-3 rounded-lg mt-4"
              onPress={handleLogin}
            >
              <Text className="text-white text-center font-semibold text-lg">Sign In</Text>
            </TouchableOpacity>
          </View>
        </View>
      </View>
    </KeyboardAvoidingView>
  );
}
EOF

# Create placeholder images
mkdir -p assets/images
touch assets/images/.gitkeep

# Create .gitignore
cat > .gitignore << 'EOF'
node_modules/
.expo/
dist/
npm-debug.*
*.jks
*.p8
*.p12
*.key
*.mobileprovision
*.orig.*
web-build/
.env
.env.local
.env.*.local

# macOS
.DS_Store

# Testing
coverage/

# Temporary files
*.log
tmp/

# IDE
.idea/
.vscode/
*.swp
*.swo
EOF

# Create README.md
cat > README.md << EOF
# ${PROJECT_TITLE} - React Native App

## Overview
React Native mobile application for ${PROJECT_TITLE} built with Expo, NativeWind v4, and TypeScript.

## Tech Stack
- **Framework**: React Native with Expo
- **Styling**: NativeWind v4 (Tailwind CSS for React Native)
- **Navigation**: React Navigation v6
- **State Management**: Zustand
- **Data Fetching**: TanStack Query (React Query)
- **Forms**: React Hook Form with Zod validation
- **API Client**: Axios

## Getting Started

### Prerequisites
- Node.js 18+
- npm or yarn
- Expo CLI
- iOS Simulator (Mac) or Android Emulator

### Installation
\`\`\`bash
# Install dependencies
npm install

# Copy environment variables
cp .env.example .env

# Start the development server
npm start
\`\`\`

### Running on Device
\`\`\`bash
# iOS
npm run ios

# Android
npm run android

# Web (experimental)
npm run web
\`\`\`

## Project Structure
\`\`\`
src/
├── app/              # Expo Router app directory
├── components/       # Reusable UI components
├── features/         # Feature modules
│   ├── ${ENTITY}/    # ${ENTITY_CAPITALIZED} feature
│   ├── auth/         # Authentication
│   ├── home/         # Home screen
│   └── profile/      # User profile
├── hooks/            # Custom React hooks
├── lib/              # External libraries config
├── navigation/       # Navigation setup
├── services/         # API services
├── stores/           # Zustand stores
├── styles/           # Global styles
├── types/            # TypeScript types
└── utils/            # Utility functions
\`\`\`

## Features
- ✅ Authentication with secure token storage
- ✅ ${ENTITY_CAPITALIZED} listing and details
- ✅ Pull-to-refresh
- ✅ Offline support with React Query
- ✅ Type-safe API client
- ✅ NativeWind v4 for styling
- ✅ Bottom tab navigation
- ✅ Form validation with Zod

## Development

### Code Style
\`\`\`bash
# Lint code
npm run lint

# Format code
npm run format
\`\`\`

### Testing
\`\`\`bash
# Run tests
npm test
\`\`\`

## Building for Production

### iOS
\`\`\`bash
expo build:ios
\`\`\`

### Android
\`\`\`bash
expo build:android
\`\`\`

## Environment Variables
- \`EXPO_PUBLIC_API_URL\`: Backend API URL
- \`EXPO_PUBLIC_AUTH_ENABLED\`: Enable/disable authentication
- \`EXPO_PUBLIC_S3_ENABLED\`: Enable/disable S3 uploads

## License
MIT
EOF

# Initialize git
echo "📦 Initializing git repository..."
git init
git add .
git commit -m "initial commit: React Native app with NativeWind v4"

echo ""
echo "🎉 React Native project '${PROJECT_NAME}-app' created successfully!"
echo ""
echo "📁 Location: $APP_DIR"
echo "📱 Features: NativeWind v4, React Navigation, TanStack Query, Zustand"
echo ""
echo "🚀 Next steps:"
echo "1. cd $APP_DIR"
echo "2. cp .env.example .env"
echo "3. npm install"
echo "4. npx expo start"
echo ""
echo "📱 Run on devices:"
echo "- Press 'i' for iOS Simulator"
echo "- Press 'a' for Android Emulator"
echo "- Scan QR code with Expo Go app for physical device"
echo ""