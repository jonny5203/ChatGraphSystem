# GenBI - Distributed Enterprise Analytics Platform

GenBI is a microservices-based business intelligence platform that enables natural language querying across distributed, heterogeneous data sources. Built with Domain-Driven Design principles, it provides federated query execution, AI-powered insights, and multi-tenant isolation with dedicated compute clusters per tenant.

## Architecture Overview

The platform consists of 19 microservices organized into 6 domain boundaries, using protocol-specific gateways to eliminate single points of failure and ensure horizontal scalability.

### Services by Domain

#### Query Intelligence Domain
- **NLP Processing Service** (Python/FastAPI) - Natural language understanding and intent extraction
- **SQL Generation Service** (Java/Spring Boot) - Converts intents to optimized SQL queries
- **Query Validation Service** (Java/Spring Boot) - Security and semantic validation

#### Data Federation Domain
- **SQL Gateway Service** (Java/Spring Boot) - JDBC connections for PostgreSQL, MySQL, Oracle, SQL Server
- **Snowflake Gateway Service** (Go) - Native Snowflake SDK integration
- **Document Gateway Service** (Go) - MongoDB wire protocol handler
- **Stream Gateway Service** (Go) - Kafka/Kinesis stream processing
- **File Gateway Service** (Go) - S3/HDFS file access (Parquet, ORC, CSV)
- **Federation Orchestrator Service** (Java/Spring Boot) - Cross-gateway query coordination
- **Query Router Service** (Java/Spring Boot) - Routes to tenant-specific Trino clusters

#### Semantic Modeling Domain
- **Model Registry Service** (Java/Spring Boot) - MDL storage and versioning
- **Model Distribution Service** (Go) - Cache replication across nodes
- **Schema Discovery Service** (Java/Spring Boot) - Auto-discovery and synchronization

#### Insight Generation Domain
- **Analytics Engine Service** (Python/FastAPI) - AI insights and anomaly detection
- **Visualization Service** (Python/FastAPI) - Dynamic chart generation

#### Tenant Management Domain
- **Tenant Orchestrator Service** (Go) - Manages isolated Trino clusters
- **Security & Access Service** (Java/Spring Boot) - RBAC, data masking, OAuth

#### Scheduling & Automation Domain
- **Workflow Engine Service** (Java/Spring Boot) - Multi-step workflow orchestration
- **Report Scheduler Service** (Java/Spring Boot) - Automated report delivery

## System Architecture
mermaidgraph TB
    subgraph "Client Layer"
        UI[Web UI]
        API[REST API]
    end

    subgraph "Query Intelligence"
        NLP[NLP Processing<br/>Python]
        SQL[SQL Generation<br/>Java]
        VAL[Query Validation<br/>Java]
    end

    subgraph "Data Federation"
        FO[Federation<br/>Orchestrator<br/>Java]
        QR[Query Router<br/>Java]
        
        subgraph "Protocol Gateways"
            SG[SQL Gateway<br/>Java]
            SF[Snowflake Gateway<br/>Go]
            DG[Document Gateway<br/>Go]
            ST[Stream Gateway<br/>Go]
            FG[File Gateway<br/>Go]
        end
    end

    subgraph "Semantic Layer"
        MR[Model Registry<br/>Java]
        MD[Model Distribution<br/>Go]
        SD[Schema Discovery<br/>Java]
    end

    subgraph "Insights"
        AE[Analytics Engine<br/>Python]
        VIZ[Visualization<br/>Python]
    end

    subgraph "Tenant Management"
        TO[Tenant Orchestrator<br/>Go]
        SEC[Security Service<br/>Java]
    end

    subgraph "Automation"
        WF[Workflow Engine<br/>Java]
        RS[Report Scheduler<br/>Java]
    end

    subgraph "Infrastructure"
        K8S[Kubernetes Clusters]
        TRINO[Trino<br/>Per Tenant]
        KAFKA[Kafka]
        REDIS[Redis]
        PG[(PostgreSQL)]
    end

    subgraph "Data Sources"
        RDB[(SQL DBs)]
        SNF[(Snowflake)]
        MDB[(MongoDB)]
        S3[S3/HDFS]
        KFK[Kafka Streams]
    end

    UI --> API
    API --> NLP
    NLP --> SQL
    SQL --> VAL
    VAL --> FO
    FO --> QR
    QR --> TRINO
    FO --> SG
    FO --> SF
    FO --> DG
    FO --> ST
    FO --> FG
    
    SG --> RDB
    SF --> SNF
    DG --> MDB
    ST --> KFK
    FG --> S3
    
    SQL -.-> MR
    VAL -.-> SEC
    FO --> AE
    AE --> VIZ
    
    TO --> K8S
    TO --> TRINO
    
    MD --> REDIS
    MR --> PG
    WF --> KAFKA
    RS --> KAFKA
    
    style NLP fill:#e1f5fe
    style SQL fill:#fff3e0
    style VAL fill:#fff3e0
    style FO fill:#fff3e0
    style QR fill:#fff3e0
    style SG fill:#fff3e0
    style SF fill:#f3e5f5
    style DG fill:#f3e5f5
    style ST fill:#f3e5f5
    style FG fill:#f3e5f5
    style MR fill:#fff3e0
    style MD fill:#f3e5f5
    style SD fill:#fff3e0
    style AE fill:#e1f5fe
    style VIZ fill:#e1f5fe
    style TO fill:#f3e5f5
    style SEC fill:#fff3e0
    style WF fill:#fff3e0
    style RS fill:#fff3e0
markdown## Key Features

- **Natural Language Querying**: Convert plain English to optimized SQL
- **Federated Query Execution**: Query multiple databases without data movement
- **Protocol-Specific Gateways**: Optimized connectors for each data source type
- **Multi-Tenant Isolation**: Dedicated Trino clusters per tenant
- **AI-Powered Insights**: Automatic anomaly detection and forecasting
- **Enterprise Security**: RBAC, data masking, GDPR/HIPAA compliance
- **Horizontal Scalability**: All services support auto-scaling
- **CQRS Architecture**: Optimized read/write paths with event sourcing

## Technology Stack

- **Languages**: Java 23.0.2, Go 1.25.0, Python 3.12.11
- **Frameworks**: Spring Boot 3.5.5, FastAPI 0.116.1
- **Query Engine**: Apache Trino
- **Orchestration**: Kubernetes 1.28, Istio Service Mesh
- **Messaging**: Apache Kafka 3.5
- **Caching**: Redis 7
- **Monitoring**: Prometheus, Grafana, Jaeger
