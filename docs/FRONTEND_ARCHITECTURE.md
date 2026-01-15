# Frontend Architecture - Next.js Template

## Overview

The frontend template is a production-ready Next.js 15 application that pairs with the Go DDD API template. It follows modern React best practices with TypeScript, featuring a clean separation of concerns and efficient state management.

## Tech Stack

### Core Framework
- **Next.js 15** - React framework with App Router for file-based routing
- **React 19** - Latest React with Server Components support
- **TypeScript** - Full type safety across the application

### State Management
- **TanStack Query v5** - Server state management, caching, and synchronization
- **Zustand** - Client state management for auth and UI state
- **React Hook Form** - Form state management with validation

### Styling & UI
- **Tailwind CSS v3** - Utility-first CSS framework
- **shadcn/ui** - Copy-paste component library built on Radix UI
- **Lucide React** - Icon library

### Data Fetching
- **Axios** - HTTP client with interceptors for auth and error handling
- **TanStack Query** - Declarative data fetching with automatic caching

## Project Structure

```
src/
├── app/                    # Next.js App Router
│   ├── layout.tsx         # Root layout with providers
│   ├── page.tsx           # Home page
│   └── [entity]/          # Dynamic entity routes
│       ├── page.tsx       # Entity list page
│       └── [id]/          # Entity detail routes
│           └── page.tsx   # Entity detail page
│
├── components/            # Reusable components
│   ├── providers.tsx      # React Query provider
│   └── ui/               # shadcn/ui components
│       ├── button.tsx
│       ├── card.tsx
│       ├── toast.tsx
│       └── ...
│
├── features/             # Feature-based modules
│   └── [entity]/        # Entity feature module
│       ├── components/  # Entity-specific components
│       │   ├── entity-list.tsx
│       │   ├── entity-form.tsx
│       │   └── entity-detail.tsx
│       ├── hooks/       # Custom hooks
│       │   └── use-entity.ts
│       ├── services/    # API service layer
│       │   └── api.ts
│       └── types/       # TypeScript types
│           └── index.ts
│
├── lib/                  # Utilities and configuration
│   ├── api/             # API client setup
│   │   ├── client.ts    # Axios instance
│   │   └── endpoints.ts # API endpoint constants
│   └── utils.ts         # Utility functions
│
├── stores/              # Zustand stores
│   ├── auth.store.ts    # Authentication state
│   └── ui.store.ts      # UI state (theme, sidebar, etc.)
│
└── types/               # Global TypeScript types
    └── index.ts
```

## Key Design Patterns

### 1. Feature-Based Architecture
Each feature (entity) is self-contained with its own:
- Components
- Hooks
- Services
- Types

This makes the codebase scalable and maintainable.

### 2. Separation of Concerns

#### API Layer (services/)
Handles all HTTP requests using Axios:
```typescript
export const entityService = {
  getAll: async () => apiClient.get('/entities'),
  getById: async (id) => apiClient.get(`/entities/${id}`),
  create: async (data) => apiClient.post('/entities', data),
  update: async (id, data) => apiClient.put(`/entities/${id}`, data),
  delete: async (id) => apiClient.delete(`/entities/${id}`)
};
```

#### Data Layer (hooks/)
Manages data fetching and caching with TanStack Query:
```typescript
export const useEntityList = () => {
  return useQuery({
    queryKey: ['entities'],
    queryFn: entityService.getAll
  });
};
```

#### Presentation Layer (components/)
React components that consume hooks and render UI.

### 3. Type Safety
All API responses and requests are typed:
```typescript
interface Entity {
  id: string;
  name: string;
  // ... other fields
}
```

### 4. Error Handling
Centralized error handling through:
- Axios interceptors for API errors
- TanStack Query error states
- Toast notifications for user feedback

## State Management Strategy

### Server State (TanStack Query)
- API data caching
- Background refetching
- Optimistic updates
- Pagination support

### Client State (Zustand)
- Authentication state
- User preferences
- UI state (theme, sidebar)
- Form drafts

### Form State (React Hook Form)
- Field validation with Zod
- Error handling
- Submission state

## Authentication Flow

1. **Login Request** → API returns JWT token
2. **Token Storage** → Zustand persist middleware stores in localStorage
3. **Request Interceptor** → Axios adds token to all requests
4. **401 Response** → Automatic logout and redirect

## Performance Optimizations

### Code Splitting
- Route-based splitting with Next.js
- Component lazy loading where appropriate

### Data Fetching
- Prefetching on hover/focus
- Stale-while-revalidate strategy
- Infinite queries for lists

### Rendering
- React Server Components for static content
- Client components only where needed
- Memoization for expensive computations

## Development Workflow

### Adding a New Feature

1. **Create Feature Module**
```bash
mkdir -p src/features/new-feature/{components,hooks,services,types}
```

2. **Define Types**
```typescript
// src/features/new-feature/types/index.ts
export interface NewFeature {
  id: string;
  // fields...
}
```

3. **Create Service**
```typescript
// src/features/new-feature/services/api.ts
export const newFeatureService = {
  // CRUD operations
};
```

4. **Create Hooks**
```typescript
// src/features/new-feature/hooks/use-new-feature.ts
export const useNewFeatureList = () => {
  return useQuery({
    queryKey: ['new-features'],
    queryFn: newFeatureService.getAll
  });
};
```

5. **Build Components**
Create UI components using shadcn/ui primitives.

## Configuration

### Environment Variables
```env
NEXT_PUBLIC_API_URL=http://localhost:8000
NEXT_PUBLIC_APP_NAME=My App
NEXT_PUBLIC_ENABLE_AUTH=true
```

### API Client Configuration
- Base URL from environment
- 10 second timeout
- Auth token injection
- Error transformation

## Testing Strategy

### Unit Tests
- Component testing with React Testing Library
- Hook testing with renderHook
- Service mocking with MSW

### Integration Tests
- API integration with real backend
- Form submission flows
- Authentication flows

### E2E Tests
- Critical user journeys with Playwright
- Cross-browser testing
- Mobile responsive testing

## Deployment Considerations

### Build Optimization
```bash
npm run build
```
- Minification
- Tree shaking
- Image optimization

### Environment Configuration
- Separate `.env` files for dev/staging/production
- API URL configuration
- Feature flags

### Docker Support
```dockerfile
FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build
CMD ["npm", "start"]
```

## Best Practices

1. **Keep components small and focused**
2. **Use TypeScript strictly**
3. **Handle loading and error states**
4. **Implement proper SEO with metadata**
5. **Follow accessibility guidelines**
6. **Use semantic HTML**
7. **Optimize images and fonts**
8. **Implement proper caching strategies**
9. **Use environment variables for configuration**
10. **Keep dependencies up to date**

## Common Patterns

### Data Table with Pagination
```typescript
const { data, isLoading } = useEntityList(page, pageSize);
```

### Form with Validation
```typescript
const form = useForm<FormData>({
  resolver: zodResolver(schema)
});
```

### Protected Routes
```typescript
const { isAuthenticated } = useAuthStore();
if (!isAuthenticated) redirect('/login');
```

### Optimistic Updates
```typescript
const mutation = useMutation({
  mutationFn: entityService.update,
  onMutate: async (newData) => {
    // Optimistically update cache
  }
});
```

This architecture ensures a scalable, maintainable, and performant frontend application that seamlessly integrates with the Go backend.