try:
    import psycopg2
except ImportError:
    print("Fatal: failed to find psycopg2")
    exit(1)
