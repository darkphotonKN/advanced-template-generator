#!/bin/bash

# Frontend Project Generator Script
# Usage: ./scripts/generate-frontend.sh project-name entity api-port frontend-port [auth] [s3]

PROJECT_NAME="$1"
ENTITY="$2"
API_PORT="$3"
FRONTEND_PORT="$4"
INCLUDE_AUTH="${5:-true}"
INCLUDE_S3="${6:-false}"

if [ -z "$PROJECT_NAME" ] || [ -z "$ENTITY" ] || [ -z "$API_PORT" ] || [ -z "$FRONTEND_PORT" ]; then
    echo "Usage: $0 <project-name> <entity> <api-port> <frontend-port> [auth] [s3]"
    echo "Example: $0 todo-app task 8050 3050 true false"
    exit 1
fi

# Frontend is created as sibling to backend in the same app folder
FRONTEND_DIR="../${PROJECT_NAME}/${PROJECT_NAME}-client"
ENTITY_PLURAL="${ENTITY}s"
ENTITY_CAPITALIZED="$(echo ${ENTITY:0:1} | tr '[:lower:]' '[:upper:]')${ENTITY:1}"
PROJECT_TITLE="$(echo ${PROJECT_NAME} | sed 's/-/ /g' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')"

echo "🚀 Generating frontend project..."
echo "- Project: $PROJECT_NAME"
echo "- Entity: $ENTITY ($ENTITY_PLURAL, $ENTITY_CAPITALIZED)"
echo "- API Port: $API_PORT"
echo "- Frontend Port: $FRONTEND_PORT"
echo "- Auth: $INCLUDE_AUTH"
echo "- S3: $INCLUDE_S3"

# Check if we're in the generator directory
if [ ! -d "../templates/nextjs-frontend" ]; then
    echo "❌ Error: Must run from generator/ directory"
    echo "Current directory: $(pwd)"
    exit 1
fi

# Copy frontend template
echo "📁 Copying frontend template..."
cp -r ../templates/nextjs-frontend/ "$FRONTEND_DIR/"

if [ $? -ne 0 ]; then
    echo "❌ Error: Failed to copy frontend template"
    exit 1
fi

cd "$FRONTEND_DIR"

# Process template files
echo "🔧 Processing template files..."

# Process package.json.tmpl
if [ -f "package.json.tmpl" ]; then
    sed "s/{{\.ProjectName}}/$PROJECT_NAME/g" package.json.tmpl | \
    sed "s/{{\.FrontendPort}}/$FRONTEND_PORT/g" > package.json
    rm package.json.tmpl
    echo "✅ Processed package.json"
fi

# Process .env.example.tmpl
if [ -f ".env.example.tmpl" ]; then
    sed "s/{{\.APIPort}}/$API_PORT/g" .env.example.tmpl | \
    sed "s/{{\.ProjectTitle}}/$PROJECT_TITLE/g" | \
    sed "s/{{\.IncludeAuth}}/$INCLUDE_AUTH/g" | \
    sed "s/{{\.IncludeS3}}/$INCLUDE_S3/g" > .env.example
    rm .env.example.tmpl
    echo "✅ Processed .env.example"
fi

# Process CLAUDE.md.tmpl
if [ -f "CLAUDE.md.tmpl" ]; then
    sed "s/{{\.ProjectTitle}}/$PROJECT_TITLE/g" CLAUDE.md.tmpl | \
    sed "s/{{\.APIPort}}/$API_PORT/g" | \
    sed "s/{{\.FrontendPort}}/$FRONTEND_PORT/g" > CLAUDE.md
    rm CLAUDE.md.tmpl
    echo "✅ Processed CLAUDE.md"
fi

# Rename entity files and directories
echo "🏗️ Renaming entity references..."

# Rename directories
if [ -d "src/features/item" ]; then
    mv "src/features/item" "src/features/$ENTITY"
    echo "✅ Renamed src/features/item to src/features/$ENTITY"
fi

if [ -d "src/app/items" ]; then
    mv "src/app/items" "src/app/$ENTITY_PLURAL"
    echo "✅ Renamed src/app/items to src/app/$ENTITY_PLURAL"
fi

# Rename files
find . -name "*item*" -type f | while read file; do
    newfile=$(echo "$file" | sed "s/item/$ENTITY/g")
    if [ "$file" != "$newfile" ]; then
        mv "$file" "$newfile"
        echo "✅ Renamed $file to $newfile"
    fi
done

# Replace content in files
echo "📝 Updating entity references in code..."

# Replace in TypeScript/JavaScript files
find . -name "*.ts" -o -name "*.tsx" -o -name "*.js" -o -name "*.jsx" | xargs sed -i '' \
    -e "s/item/$ENTITY/g" \
    -e "s/Item/$ENTITY_CAPITALIZED/g" \
    -e "s/items/$ENTITY_PLURAL/g" \
    -e "s/Items/$ENTITY_CAPITALIZED$ENTITY_PLURAL/g"

echo "✅ Updated TypeScript files"

# Initialize git
echo "📦 Initializing git repository..."
git init
git add .
git commit -m "initial commit"

echo ""
echo "🎉 Frontend project '$PROJECT_NAME-client' created successfully!"
echo ""
echo "📁 Location: $FRONTEND_DIR"
echo "🌐 Frontend will run on: http://localhost:$FRONTEND_PORT"
echo "🔗 Connected to API on: http://localhost:$API_PORT"
echo ""
echo "🚀 Next steps:"
echo "1. cd $FRONTEND_DIR"
echo "2. cp .env.example .env"
echo "3. npm install"
echo "4. npm run dev"
echo ""